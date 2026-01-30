package enum

// StatusCodeOrder adalah hasil eksekusi API (success / business error).
type StatusCodeOrder string

const (
	// Validation / Business Errors
	ErrTotalPriceMismatch    StatusCodeOrder = "ERR_TOTAL_PRICE_MISMATCH"
	ErrInvalidPaymentMethod  StatusCodeOrder = "ERR_INVALID_PAYMENT_METHOD"
	ErrFailedChargePayment   StatusCodeOrder = "ERR_FAILED_CHARGE_PAYMENT"
	ErrPlatformFeeIsZero     StatusCodeOrder = "ERR_PLATFORM_FEE_IS_ZERO"
	ErrOutOfStock            StatusCodeOrder = "ERR_OUT_OF_STOCK"
	ErrInvalidShippingMethod StatusCodeOrder = "ERR_INVALID_SHIPPING_METHOD"

	// Success
	SuccessOrderCreated StatusCodeOrder = "SUCCESS_ORDER_CREATED"
)

// OrderStatus adalah lifecycle pesanan di e-commerce.
type OrderStatus string

const (
	// Order dibuat, menunggu pembayaran
	OrderStatusPendingPayment OrderStatus = "PENDING_PAYMENT"

	// Pembayaran berhasil, menunggu diproses seller
	OrderStatusPaid OrderStatus = "PAID"

	// Order sedang diproses / dikemas
	OrderStatusProcessing OrderStatus = "PROCESSING"

	// Order dikirim
	OrderStatusShipped OrderStatus = "SHIPPED"

	// Order telah diterima customer
	OrderStatusDelivered OrderStatus = "DELIVERED"

	// Order selesai (complete)
	OrderStatusCompleted OrderStatus = "COMPLETED"

	// Order dibatalkan
	OrderStatusCanceled OrderStatus = "CANCELED"

	// Order gagal (system / payment error)
	OrderStatusFailed OrderStatus = "FAILED"
)

// OrderPaymentStatus adalah lifecycle pembayaran order.
type OrderPaymentStatus string

const (
	// Menunggu pembayaran
	PaymentStatusPending OrderPaymentStatus = "PENDING"

	// Pembayaran berhasil
	PaymentStatusPaid OrderPaymentStatus = "PAID"

	// Pembayaran gagal
	PaymentStatusFailed OrderPaymentStatus = "FAILED"

	// Pembayaran expired
	PaymentStatusExpired OrderPaymentStatus = "EXPIRED"

	// Pembayaran dibatalkan
	PaymentStatusCanceled OrderPaymentStatus = "CANCELED"

	// Refund sedang diproses
	PaymentStatusRefunding OrderPaymentStatus = "REFUNDING"

	// Refund selesai
	PaymentStatusRefunded OrderPaymentStatus = "REFUNDED"
)
