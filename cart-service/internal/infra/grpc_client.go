package infra

import (
	pb "cart-service/proto/product"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewProductClient() (pb.ProductServiceClient, func(), error) {
	conn, err := grpc.NewClient(
        "localhost:50051",
        grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
    if err != nil {
        slog.Error("did not connect", "error", err)
        return nil, nil, err
    }

    client := pb.NewProductServiceClient(conn)

    cleanup := func() {
        if err := conn.Close(); err != nil {
            slog.Error("failed to close grpc connection", "error", err)
        }
    }

	return client, cleanup, nil
}

