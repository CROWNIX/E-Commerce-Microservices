package models

type Cart struct {
	ID        uint64 `db:"id" json:"id"`
	UserID    uint64 `db:"user_id" json:"user_id"`
	ProductID uint64 `db:"product_id" json:"product_id"`
	Quantity  uint8  `db:"quantity" json:"quantity"`

	Audit
}

const CartTableName = "carts"

var CartField = struct {
	ID        string
	UserID    string
	ProductID string
	Quantity  string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	ProductID: "product_id",
	Quantity:  "quantity",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}