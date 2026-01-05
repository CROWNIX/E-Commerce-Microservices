package infra

import (
	"pkg/client/product"
	pb "pkg/proto/generated/product"
)

func NewProductClient() (pb.ProductServiceClient, func(), error) {
	// TODO: Get target from config
	return product.NewProductClient("localhost:50051")
}
