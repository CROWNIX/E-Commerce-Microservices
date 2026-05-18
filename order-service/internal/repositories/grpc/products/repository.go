package products

import (
	"context"
	pb "pkg/proto/generated/product"

	productServiceDto "pkg/services/product-service/dto"

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

func (r *productRepository) GetDetailProduct(ctx context.Context, productID uint64) (output productServiceDto.GetDetailProductOutput, err error) {
	resp, err := r.client.GetProductDetail(ctx, &pb.GetProductDetailRequest{
		ProductId: productID,
	})
	if err != nil {
		return output, err
	}

	output = productServiceDto.GetDetailProductOutput{
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

func (r *productRepository) CountProductByIds(ctx context.Context, productIds []uint64) (uint32, error) {
	resp, err := r.client.CountProductByIds(ctx, &pb.CountProductByIdsRequest{
		ProductIds: productIds,
	})
	if err != nil {
		return 0, err
	}
	return resp.Total, nil
}

func (r *productRepository) GetProductByIds(ctx context.Context, productIds []uint64) ([]productServiceDto.GetDetailProductOutput, error) {
	resp, err := r.client.GetProductByIds(ctx, &pb.GetProductByIdsRequest{
		ProductIds: productIds,
	})

	if err != nil {
		return nil, err
	}

	output := make([]productServiceDto.GetDetailProductOutput, 0, len(resp.Products))
	for _, p := range resp.Products {
		dto := productServiceDto.GetDetailProductOutput{
			ID:              p.Id,
			Name:            p.Name,
			Description:     p.Description,
			Price:           p.Price,
			Stock:           p.Stock,
			FinalPrice:      p.FinalPrice,
			DiscountPercent: uint8(p.DiscountPercent),
			MinimumPurchase: uint8(p.MinimumPurchase),
		}

		if len(p.Images) > 0 {
			dto.Images = databases.NewJSON(p.Images)
		} else {
			dto.Images = databases.NewJSON([]string{})
		}

		if p.MaximumPurchase != nil {
			mp := uint8(*p.MaximumPurchase)
			dto.MaximumPurchase = &mp
		}

		output = append(output, dto)
	}

	return output, nil
}
