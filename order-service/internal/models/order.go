package models

type Order struct {
	ID              uint64  `db:"id" json:"id"`
	UserID          uint64  `db:"user_id" json:"user_id"`
	AddressID       uint64  `db:"address_id" json:"address_id"`
	PaymentMethodID uint64  `db:"payment_method_id" json:"payment_method_id"`
	Status          string  `db:"status" json:"status"`
	PaymentStatus   *string `db:"payment_status" json:"payment_status"`

	Audit
}

const OrderTableName = "orders"

var OrderField = struct {
	ID              string
	UserID          string
	AddressID       string
	PaymentMethodID string
	Status          string
	PaymentStatus   string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
}{
	ID:              "id",
	UserID:          "user_id",
	AddressID:       "address_id",
	PaymentMethodID: "payment_method_id",
	Status:          "status",
	PaymentStatus:   "payment_status",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}
