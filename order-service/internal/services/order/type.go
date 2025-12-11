package order

import "encoding/json"

type PaymentInstruction interface {
    Payload() map[string]any
    Bytes() ([]byte, error)
}

type Payment[T any] struct {
	Code string
	Data T
}

func (p Payment[T]) Payload() map[string]any {
	payload := map[string]any{
		"payment_code": p.Code,
	}

	b, _ := json.Marshal(p.Data)
	_ = json.Unmarshal(b, &payload)

	return payload
}

func (p Payment[T]) Bytes() ([]byte, error) {
	return json.Marshal(p.Payload())
}

type CreateOrderServiceOutput struct{
	Status string
	PaymentInstruction PaymentInstruction
}

type CreateOrderServiceInput struct {
	UserID          uint64
	AddressID       uint64
	PaymentMethodID uint64
	GrandTotal      uint64
	Items           []Item
}

type Item struct {
	ProductId uint64
	Quantity  uint64
}

type QrisData struct {
    QrisContent   string `json:"qris_content"`
    QrisURL       string `json:"qris_url"`
    GenerateQrUrl string `json:"generate_qr_url"`
}
