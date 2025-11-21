package carts

type CreateCartInput struct {
	UserID    uint64
	ProductID uint64
	Quantity  uint8
}