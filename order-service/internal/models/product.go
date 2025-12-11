package models

import "github.com/CROWNIX/go-utils/databases"

type Product struct {
	ID              uint64                   `db:"id"`
	CategoryID      string                   `db:"category_id"`
	Name            string                   `db:"name"`
	Images          databases.JSON[[]string] `db:"images"`
	Description     string                   `db:"description"`
	Price           uint64                   `db:"price"`
	Stock           uint32                   `db:"stock"`
	FinalPrice      uint64                   `db:"final_price"`
	DiscountPercent uint8                    `db:"discount_percent"`
	MinimumPurchase uint8                    `db:"minimum_purchase"`
	MaximumPurchase *uint8                   `db:"maximum_purchase"`
	Audit
}

const ProductTableName = "products"

var ProductField = struct {
	ID              string
	CategoryID      string
	Name            string
	Images          string
	Description     string
	Price           string
	Stock           string
	FinalPrice      string
	DiscountPercent string
	MinimumPurchase string
	MaximumPurchase string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
}{
	ID:              "id",
	CategoryID:      "category_id",
	Name:            "name",
	Images:          "images",
	Description:     "description",
	Price:           "price",
	Stock:           "stock",
	FinalPrice:      "final_price",
	DiscountPercent: "discount_percent",
	MinimumPurchase: "minimum_purchase",
	MaximumPurchase: "maximum_purchase",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}
