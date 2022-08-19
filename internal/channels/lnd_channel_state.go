package channels

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
)

type lndGetLatestChannelsStateConnection interface {
	GetChanInfo(ctx context.Context, in *lnrpc.ChanInfoRequest, opts ...grpc.CallOption) (*lnrpc.ChannelEdge, error)
	ListChannels(ctx context.Context, in *lnrpc.ListChannelsRequest, opts ...grpc.CallOption) (*lnrpc.ListChannelsResponse, error)
}

func GetLatestChannelsState(client lndGetLatestChannelsStateConnection, db *sqlx.DB, localNodeId int) {
	client.GetChanInfo(context.Background(), &lnrpc.ChanInfoRequest{})
}
