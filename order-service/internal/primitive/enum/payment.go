package enum

type PaymentType string

const (
	PMQris PaymentType = "qris"
	PMCash PaymentType = "cash"
)
