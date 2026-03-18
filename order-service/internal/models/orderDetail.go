package models

type OrderDetail struct {
	OrderID       uint64  `db:"order_id" json:"order_id"`
	ProductID     uint64  `db:"product_id" json:"product_id"`
	Quantity      uint8   `db:"payment_method_id" json:"payment_method_id"`
	Status        string  `db:"status" json:"status"`
	PaymentStatus *string `db:"payment_status" json:"payment_status"`
}

const OrderDetailTableName = "order_details"

var OrderDetailField = struct {
	OrderID       string
	ProductID     string
	Quantity      string
	Status        string
	PaymentStatus string
}{
	OrderID:       "id",
	ProductID:     "product_id",
	Quantity:      "quantity",
	Status:        "status",
	PaymentStatus: "payment_status",
}
