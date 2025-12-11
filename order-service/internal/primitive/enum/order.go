package enum

// StatusCodeOrder adalah hasil eksekusi API (success/error).
type StatusCodeOrder string

const (
	ErrTotalPriceMismatch StatusCodeOrder = "ERR_TOTAL_PRICE_MISMATCH"

	ErrInvalidPaymentMethod StatusCodeOrder = "ERR_INVALID_PAYMENT_METHOD"

	ErrFailedChargePayment StatusCodeOrder = "ERR_FAILED_CHARGE_PAYMENT"

	ErrPlatformFeeIsZero StatusCodeOrder = "ERR_PLATFORM_FEE_IS_ZERO"

	// Success: order created successfully
	SuccessOrder StatusCodeOrder = "SUCCESS"
)

// OrderStatus adalah state lifecycle dari Order di sistem.
type OrderStatus string

const (
	// Order baru dibuat, menunggu pembayaran
	OrderStatusOnOrder OrderStatus = "ON_ORDER"
	// Order sedang diproses / dilayani setelah pembayaran
	OrderStatusIsBeingServed OrderStatus = "IS_BEING_SERVED"
	// Order dibatalkan (oleh user/system)
	OrderStatusCanceled OrderStatus = "ORDER_CANCELED"
	// Order gagal (contoh: payment error, expired)
	OrderStatusFailed OrderStatus = "ORDER_FAILED"
)

// OrderPaymentStatus adalah state lifecycle dari pembayaran.
type OrderPaymentStatus string

const (
	// Baru dibuat, menunggu pembayaran
	PaymentStatusWaiting OrderPaymentStatus = "WAITING_PAYMENT"
	// Pembayaran berhasil
	PaymentStatusCompleted OrderPaymentStatus = "PAYMENT_COMPLETED"
	// Pembayaran expired (habis waktu)
	PaymentStatusExpired OrderPaymentStatus = "PAYMENT_EXPIRED"
	// Pembayaran dibatalkan
	PaymentStatusCanceled OrderPaymentStatus = "PAYMENT_CANCELED"
	// Pembayaran gagal (error)
	PaymentStatusFailed OrderPaymentStatus = "PAYMENT_FAILED"
)
