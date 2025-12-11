package payment

import (
	"context"
)

type CoreApi interface {
	Charge(ctx context.Context, input ChargeRequest) (output ChargeOutput, err error)
	GetPaymentStatusMidtrans(ctx context.Context, id uint64) (output PaymentStatusOutput, err error)
}