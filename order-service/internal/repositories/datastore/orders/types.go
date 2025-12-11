package orders

type CreateOrderInput struct {
	UserID          uint64
	AddressID       uint64
	PaymentMethodID uint64
	GrandTotal      uint64
	Status          string
	PaymentStatus   string
	// Items           []Item
}

type Item struct {
	ProductId uint64
	Quantity  uint64
}
