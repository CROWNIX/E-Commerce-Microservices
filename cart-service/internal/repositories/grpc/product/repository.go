package product

import (
	pb "pkg/proto/product/generated"
	"context"

	"github.com/CROWNIX/go-utils/databases"
)

type productRepository struct {
	client pb.ProductServiceClient
}

func NewProductRepository(client pb.ProductServiceClient) ProductServiceInterfaces {
	return &productRepository{
		client: client,
	}
}

func (r *productRepository) GetDetailProduct(ctx context.Context, productID uint64) (output GetDetailProductOutput, err error) {
	resp, err := r.client.GetProductDetail(ctx, &pb.GetProductDetailRequest{
		ProductId: productID,
	})
	if err != nil {
		return output, err
	}

	output = GetDetailProductOutput{
		ID:              resp.Id,
		Name:            resp.Name,
		Description:     resp.Description,
		Price:           resp.Price,
		Stock:           resp.Stock,
		FinalPrice:      resp.FinalPrice,
		DiscountPercent: uint8(resp.DiscountPercent),
		MinimumPurchase: uint8(resp.MinimumPurchase),
	}

	if len(resp.Images) > 0 {
		output.Images = databases.NewJSON(resp.Images)
	} else {
		output.Images = databases.NewJSON([]string{})
	}

	if resp.MaximumPurchase != nil {
		mp := uint8(*resp.MaximumPurchase)
		output.MaximumPurchase = &mp
	}

	return output, nil
}
