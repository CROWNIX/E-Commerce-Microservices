package grpc

import (
	"context"
	pb "pkg/proto/generated/product"
	"product-service/internal/repositories/datastore/products"
	"product-service/internal/services/product"
)

type Server struct {
	pb.UnimplementedProductServiceServer
	productService product.ProductServiceInterfaces
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

func (s *Server) CountProductByIds(ctx context.Context, request *pb.CountProductByIdsRequest) (*pb.CountProductByIdsResponse, error){
	total, err := s.productRepositoryReader.CountProductByIds(ctx, request.ProductIds)
	if err != nil {
		return nil, err
	}

	response := &pb.CountProductByIdsResponse{
		Total: total,
	}

	return response, nil
}
