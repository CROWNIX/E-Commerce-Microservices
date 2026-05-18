package grpc

import (
	"context"
	pb "pkg/proto/generated/product"
	"product-service/internal/repositories/datastore/products"
	"product-service/internal/services/product"

	"github.com/CROWNIX/go-utils/utils/generic"
)

type Server struct {
	pb.UnimplementedProductServiceServer
	productService          product.ProductServiceInterfaces
	productRepositoryReader products.ProductRepositoryReaderInterfaces
}

func NewServer(productService product.ProductServiceInterfaces) *Server {
	return &Server{
		productService: productService,
	}
}

func (s *Server) GetProductDetail(ctx context.Context, req *pb.GetProductDetailRequest) (*pb.GetProductDetailResponse, error) {
	result, err := s.productService.GetDetailProduct(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetProductDetailResponse{
		Id:              result.ID,
		Name:            result.Name,
		Images:          result.Images.V,
		Description:     result.Description,
		Price:           result.Price,
		Stock:           result.Stock,
		FinalPrice:      result.FinalPrice,
		DiscountPercent: uint32(result.DiscountPercent),
		MinimumPurchase: uint32(result.MinimumPurchase),
	}

	if result.MaximumPurchase != nil {
		maxPurchase := uint32(*result.MaximumPurchase)
		response.MaximumPurchase = &maxPurchase
	}

	return response, nil
}

func (s *Server) CountProductByIds(ctx context.Context, request *pb.CountProductByIdsRequest) (*pb.CountProductByIdsResponse, error) {
	total, err := s.productRepositoryReader.CountProductByIds(ctx, request.ProductIds)
	if err != nil {
		return nil, err
	}

	response := &pb.CountProductByIdsResponse{
		Total: total,
	}

	return response, nil
}

func (s *Server) GetProductByIds(ctx context.Context, request *pb.GetProductByIdsRequest) (*pb.GetProductByIdsResponse, error) {
	productsOutput, err := s.productService.GetProductsByIds(ctx, request.ProductIds)
	if err != nil {
		return nil, err
	}

	response := generic.TransformSlice(productsOutput, func(product product.GetDetailProductOutput) *pb.GetProductDetailResponse {
		var maxPurchase *uint32

		if product.MaximumPurchase != nil {
			v := uint32(*product.MaximumPurchase)
			maxPurchase = &v
		}

		return &pb.GetProductDetailResponse{
			Id:              product.ID,
			Name:            product.Name,
			Images:          product.Images.V,
			Description:     product.Description,
			Price:           product.Price,
			Stock:           product.Stock,
			FinalPrice:      product.FinalPrice,
			DiscountPercent: uint32(product.DiscountPercent),
			MinimumPurchase: uint32(product.MinimumPurchase),
			MaximumPurchase: maxPurchase,
		}
	})

	output := pb.GetProductByIdsResponse{
		Products: response,
	}

	return &output, nil
}
