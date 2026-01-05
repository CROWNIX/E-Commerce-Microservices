package product

import (
	"log/slog"

	product "pkg/proto/generated/product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewProductClient(target string) (product.ProductServiceClient, func(), error) {
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		slog.Error("did not connect", "error", err)
		return nil, nil, err
	}

	client := product.NewProductServiceClient(conn)

	cleanup := func() {
		if err := conn.Close(); err != nil {
			slog.Error("failed to close grpc connection", "error", err)
		}
	}

	return client, cleanup, nil
}
