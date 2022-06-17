package lnd

import (
	"context"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/zpay32"
	"google.golang.org/grpc"
	"io"
	"time"
)

type invoicesClient interface {
	SubscribeInvoices(ctx context.Context, in *lnrpc.InvoiceSubscription,
		opts ...grpc.CallOption) (lnrpc.Lightning_SubscribeInvoicesClient, error)
}

func fetchLastInvoiceIndexes(db *sqlx.DB) (addIndex uint64, settleIndex uint64, err error) {

	addIndex = 1
	settleIndex = 1

	sqlLatest := `select max(add_index), max(settle_index) from invoice;`

	row := db.QueryRow(sqlLatest)
	err = row.Scan(&addIndex, &settleIndex)

	switch err {
	case nil:
		return addIndex, settleIndex, err
	case sql.ErrNoRows:
		return 1, 1, nil
	}

	return addIndex, settleIndex, nil
}

func SubscribeAndStoreInvoices(ctx context.Context, client invoicesClient, db *sqlx.DB) error {

	// Get the latest settle and add index to prevent duplicate entries.
	addIndex, settleIndex, err := fetchLastInvoiceIndexes(db)

	invoiceStream, err := client.SubscribeInvoices(ctx, &lnrpc.InvoiceSubscription{
		AddIndex:    addIndex,
		SettleIndex: settleIndex,
	})
	if err != nil {
		return err
	}

	for {

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		invoice, err := invoiceStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		// TODO: Check the running nodes network. Currently we assume we are running on Bitcoin mainnet
		inva, err := zpay32.Decode(invoice.PaymentRequest, &chaincfg.MainNetParams)
		if err != nil {
			return err
		}

		err = insertInvoice(db, invoice, fmt.Sprintf("%x", inva.Destination.SerializeCompressed()))
		if err != nil {
			return err
		}

	}

	return nil
}

var sqlInvoice = `INSERT INTO invoice (
    memo,
    r_preimage,
    r_hash,
    value_msat,
    creation_date,
    settle_date,
    payment_request,
    destination_pub_key,
    description_hash,
    expiry,
    fallback_addr,
    cltv_expiry,
    route_hints,
    private,
    add_index,
    settle_index,
    amt_paid_msat,
    /*
    The state the invoice is in.
        OPEN = 0;
        SETTLED = 1;
        CANCELED = 2;
        ACCEPTED = 3;
    */
    invoice_state,
    htlcs,
    features,
    is_keysend,
    payment_addr,
    is_amp,
    amp_invoice_state,
    created_on,
    updated_on
) VALUES(
	:memo,
    :r_preimage,
    :r_hash,
    :value_msat,
    :creation_date,
    :settle_date,
    :payment_request,
	:destination_pub_key,
    :description_hash,
    :expiry,
    :fallback_addr,
    :cltv_expiry,
    :route_hints,
    :private,
    :add_index,
    :settle_index,
    :amt_paid_msat,
    :invoice_state,
    :htlcs,
    :features,
    :is_keysend,
    :payment_addr,
    :is_amp,
    :amp_invoice_state,
    :created_on,
    :updated_on
);`

func insertInvoice(db *sqlx.DB, invoice *lnrpc.Invoice, destination string) error {

	rhJson, err := json.Marshal(invoice.RouteHints)
	if err != nil {
		return err
	}

	htlcJson, err := json.Marshal(invoice.Htlcs)
	if err != nil {
		return err
	}

	featuresJson, err := json.Marshal(invoice.Features)
	if err != nil {
		return err
	}

	aisJson, err := json.Marshal(invoice.AmpInvoiceState)
	if err != nil {
		return err
	}

	i := Invoice{
		Memo:            invoice.Memo,
		RPreimage:       hex.EncodeToString(invoice.RPreimage),
		RHash:           hex.EncodeToString(invoice.RHash),
		ValueMsat:       invoice.ValueMsat,
		CreationDate:    time.Unix(invoice.CreationDate, 0).UTC(),
		SettleDate:      time.Unix(invoice.SettleDate, 0).UTC(),
		PaymentRequest:  invoice.PaymentRequest,
		Destination:     destination,
		DescriptionHash: invoice.DescriptionHash,
		Expiry:          invoice.Expiry,
		FallbackAddr:    invoice.FallbackAddr,
		CltvExpiry:      invoice.CltvExpiry,
		RouteHints:      rhJson,
		Private:         false,
		AddIndex:        invoice.AddIndex,
		SettleIndex:     invoice.SettleIndex,
		AmtPaidSat:      invoice.AmtPaidSat,
		AmtPaidMsat:     invoice.AmtPaidMsat,
		InvoiceState:    invoice.State.String(), // ,
		Htlcs:           htlcJson,
		Features:        featuresJson,
		IsKeysend:       invoice.IsKeysend,
		PaymentAddr:     hex.EncodeToString(invoice.PaymentAddr),
		IsAmp:           invoice.IsAmp,
		AmpInvoiceState: aisJson,
		CreatedOn:       time.Now().UTC(),
		UpdatedOn:       nil,
	}

	_, err = db.NamedExec(sqlInvoice, i)

	if err != nil {
		return errors.Wrapf(err, `insertInvoice -> db.Exec(%s, ...variables)`, sqlInvoice)
	}
	return nil
}

type Invoice struct {

	/*
	   An optional memo to attach along with the invoice. Used for record keeping
	   purposes for the invoice's creator, and will also be set in the description
	   field of the encoded payment request if the description_hash field is not
	   being used.
	*/
	Memo string `db:"memo" json:"memo"`

	/*
	   The hex-encoded preimage (32 byte) which will allow settling an incoming
	   HTLC payable to this preimage.
	*/
	RPreimage string `db:"r_preimage" json:"r_preimage"`

	/*
	   The hash of the preimage.
	*/
	RHash string `db:"r_hash" json:"r_hash"`

	// The value of the invoice
	ValueMsat int64 `db:"value_msat" json:"value_msat"`

	// When this invoice was created
	CreationDate time.Time `db:"creation_date" json:"creation_date"`

	// When this invoice was settled
	SettleDate time.Time `db:"settle_date" json:"settle_date"`

	/*
	   A bare-bones invoice for a payment within the Lightning Network. With the
	   details of the invoice, the sender has all the data necessary to send a
	   payment to the recipient.
	*/
	PaymentRequest string `db:"payment_request" json:"payment_request"`

	/*
	   A bare-bones invoice for a payment within the Lightning Network. With the
	   details of the invoice, the sender has all the data necessary to send a
	   payment to the recipient.
	*/
	Destination string `db:"destination_pub_key" json:"destination_pub_key"`

	/*
	   Hash (SHA-256) of a description of the payment. Used if the description of
	   payment (memo) is too long to naturally fit within the description field
	   of an encoded payment request.
	*/
	DescriptionHash []byte `db:"description_hash" json:"description_hash"`

	// Payment request expiry time in seconds. Default is 3600 (1 hour).
	Expiry int64 `db:"expiry" json:"expiry"`

	// Fallback on-chain address.
	FallbackAddr string `db:"fallback_addr" json:"fallback_addr"`

	// Delta to use for the time-lock of the CLTV extended to the final hop.
	CltvExpiry uint64 `db:"cltv_expiry" json:"cltv_expiry"`

	/*
	   Route hints that can each be individually used to assist in reaching the
	   invoice's destination.
	*/
	//repeated RouteHint route_hints = 14;
	RouteHints []byte `db:"route_hints" json:"route_hints"`

	// Whether this invoice should include routing hints for private channels.
	Private bool `db:"private" json:"private"`

	/*
	   The "add" index of this invoice. Each newly created invoice will increment
	   this index making it monotonically increasing. Callers to the
	   SubscribeInvoices call can use this to instantly get notified of all added
	   invoices with an add_index greater than this one.
	*/
	AddIndex uint64 `db:"add_index" json:"add_index"`

	/*
	   The "settle" index of this invoice. Each newly settled invoice will
	   increment this index making it monotonically increasing. Callers to the
	   SubscribeInvoices call can use this to instantly get notified of all
	   settled invoices with an settle_index greater than this one.
	*/
	SettleIndex uint64 `db:"settle_index" json:"settle_index"`

	/*
	   The amount that was accepted for this invoice, in satoshis. This will ONLY
	   be set if this invoice has been settled. We provide this field as if the
	   invoice was created with a zero value, then we need to record what amount
	   was ultimately accepted. Additionally, it's possible that the sender paid
	   MORE that was specified in the original invoice. So we'll record that here
	   as well.
	*/
	AmtPaidSat int64 `db:"amt_paid_sat" json:"amt_paid_sat"`

	/*
	   The amount that was accepted for this invoice, in millisatoshis. This will
	   ONLY be set if this invoice has been settled. We provide this field as if
	   the invoice was created with a zero value, then we need to record what
	   amount was ultimately accepted. Additionally, it's possible that the sender
	   paid MORE that was specified in the original invoice. So we'll record that
	   here as well.
	*/
	AmtPaidMsat int64 `db:"amt_paid_msat" json:"amt_paid_msat"`

	InvoiceState string `db:"invoice_state" json:"invoice_state"`
	//OPEN = 0;
	//SETTLED = 1;
	//CANCELED = 2;
	//ACCEPTED = 3;

	// List of HTLCs paying to this invoice [EXPERIMENTAL].
	Htlcs []byte `db:"htlcs" json:"htlcs"`
	//repeated InvoiceHTLC htlcs = 22;

	// List of features advertised on the invoice.
	//map<uint32, Feature> features = 24;
	// features []*lnrpc.Feature
	Features []byte `db:"features" json:"features"`

	/*
	   Indicates if this invoice was a spontaneous payment that arrived via keysend
	   [EXPERIMENTAL].
	*/
	IsKeysend bool `db:"is_keysend" json:"is_keysend"`

	/*
	   The payment address of this invoice. This value will be used in MPP
	   payments, and also for newer invoices that always require the MPP payload
	   for added end-to-end security.
	*/
	PaymentAddr string `db:"payment_addr" json:"payment_addr"`

	/*
	   Signals whether this is an AMP invoice.
	*/
	IsAmp bool `db:"is_amp" json:"is_amp"`

	/*
	   [EXPERIMENTAL]:
	   Maps a 32-byte hex-encoded set ID to the sub-invoice AMP state for the
	   given set ID. This field is always populated for AMP invoices, and can be
	   used alongside LookupInvoice to obtain the HTLC information related to a
	   given sub-invoice.
	*/
	//map<string, AMPInvoiceState> amp_invoice_state = 28;
	AmpInvoiceState []byte     `db:"amp_invoice_state" json:"amp_invoice_state"`
	CreatedOn       time.Time  `db:"created_on" json:"created_on"`
	UpdatedOn       *time.Time `db:"updated_on" json:"updated_on"`
}
