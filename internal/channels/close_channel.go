package channels

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lncapital/torq/internal/settings"
	"github.com/lncapital/torq/pkg/lnd_connect"
	"google.golang.org/grpc"
	"io"
	"strconv"
	"strings"
)

type lndClientCloseChannel interface {
	CloseChannel(ctx context.Context, in *lnrpc.CloseChannelRequest, opts ...grpc.CallOption) (lnrpc.Lightning_CloseChannelClient, error)
}

type CloseChannelRequest struct {
	NodeId          int     `json:"nodeId"`
	ChannelPoint    string  `json:"channelPoint"`
	Force           *bool   `json:"force"`
	TargetConf      *int32  `json:"targetConf"`
	DeliveryAddress *string `json:"deliveryAddress"`
	SatPerVbyte     *uint64 `json:"satPerVbyte"`
}

type pendingUpdate struct {
	TxId        []byte `json:"txId"`
	OutputIndex uint32 `json:"outputIndex"`
}

type channelCloseUpdate struct {
	ClosingTxId []byte `json:"closingTxId"`
	Success     bool   `json:"success"`
}

type CloseChannelResponse struct {
	ReqId        string             `json:"reqId"`
	Status       string             `json:"status"`
	ClosePending pendingUpdate      `json:"closePending"`
	ChanClose    channelCloseUpdate `json:"chanClose"`
}

func CloseChannel(wChan chan interface{}, db *sqlx.DB, c *gin.Context, ccReq CloseChannelRequest, reqId string) (err error) {
	connectionDetails, err := settings.GetNodeConnectionDetailsById(db, ccReq.NodeId)
	if err != nil {
		return errors.New("Getting node connection details from the db")
	}

	conn, err := lnd_connect.Connect(
		connectionDetails.GRPCAddress,
		connectionDetails.TLSFileBytes,
		connectionDetails.MacaroonFileBytes)
	if err != nil {
		return errors.Wrap(err, "Connecting to LND")
	}
	defer conn.Close()
	client := lnrpc.NewLightningClient(conn)

	closeChanReq, err := prepareCloseRequest(ccReq)
	if err != nil {
		return errors.Wrap(err, "Preparing close request")
	}

	return closeChannelResp(client, &closeChanReq, wChan, reqId)
}

func prepareCloseRequest(ccReq CloseChannelRequest) (r lnrpc.CloseChannelRequest, err error) {

	if ccReq.NodeId == 0 {
		return r, errors.New("Node id is missing")
	}

	if ccReq.SatPerVbyte != nil && ccReq.TargetConf != nil {
		return r, errors.New("Cannot set both SatPerVbyte and TargetConf")
	}

	channelPoint, err := convertChannelPoint(ccReq.ChannelPoint)
	if err != nil {
		return r, err
	}
	//
	//Make the close channel request
	closeChanReq := lnrpc.CloseChannelRequest{
		ChannelPoint: channelPoint,
	}

	if ccReq.Force != nil {
		closeChanReq.Force = *ccReq.Force
	}

	if ccReq.TargetConf != nil {
		closeChanReq.TargetConf = *ccReq.TargetConf
	}

	if ccReq.SatPerVbyte != nil {
		closeChanReq.SatPerVbyte = *ccReq.SatPerVbyte
	}

	if ccReq.DeliveryAddress != nil {
		closeChanReq.DeliveryAddress = *ccReq.DeliveryAddress
	}

	return closeChanReq, nil
}

func closeChannelResp(client lndClientCloseChannel, closeChanReq *lnrpc.CloseChannelRequest, wChan chan interface{}, reqId string) error {

	ctx := context.Background()
	closeChanRes, err := client.CloseChannel(ctx, closeChanReq)
	if err != nil {
		return errors.Wrap(err, "Closing channel")
	}

	for {
		select {
		case <-ctx.Done():
			//log.Debug().Msgf("%v", ctx.Err())
			return nil
		default:
		}

		resp, err := closeChanRes.Recv()
		if err == io.EOF {
			//log.Debug().Msgf("Close channel EOF")
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "Close channel request receive")
		}

		r, err := processCloseResponse(resp, reqId)
		if err != nil {
			return errors.Wrap(err, "Process close response")
		}
		wChan <- r
	}
}

func convertChannelPoint(chanPointStr string) (chanPoint *lnrpc.ChannelPoint, err error) {
	splitChanPoint := strings.Split(chanPointStr, ":")
	if len(splitChanPoint) != 2 {
		return chanPoint, errors.New("Channel point missing a colon")
	}

	fundingTxid := &lnrpc.ChannelPoint_FundingTxidStr{FundingTxidStr: splitChanPoint[0]}

	oIndxUint, err := strconv.ParseUint(splitChanPoint[1], 10, 1)
	if err != nil {
		return chanPoint, errors.New("Parsing channel point output index")
	}

	outputIndex := uint32(oIndxUint)
	chanPoint = &lnrpc.ChannelPoint{
		FundingTxid: fundingTxid,
		OutputIndex: outputIndex,
	}

	return chanPoint, nil
}

func processCloseResponse(resp *lnrpc.CloseStatusUpdate, reqId string) (*CloseChannelResponse, error) {
	switch resp.GetUpdate().(type) {
	case *lnrpc.CloseStatusUpdate_ClosePending:
		//log.Info().Msgf("Pending closing channel")
		ccPending := resp.GetClosePending()
		return &CloseChannelResponse{
			ReqId:        reqId,
			Status:       "PENDING",
			ClosePending: pendingUpdate{ccPending.Txid, ccPending.OutputIndex},
		}, nil
	case *lnrpc.CloseStatusUpdate_ChanClose:
		//log.Info().Msgf("Channel closed/or not")
		ccClose := resp.GetChanClose()
		//if !ccClose.Success {
		//	log.Error().Msgf("Channel closing failed")
		//	return nil, errors.New("Channel failed to close")
		//}
		return &CloseChannelResponse{
			ReqId:     reqId,
			Status:    "CLOSED",
			ChanClose: channelCloseUpdate{ccClose.ClosingTxid, ccClose.Success},
		}, nil
	default:
	}

	return nil, nil
}
