package order

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"order-service/internal/config"
	"order-service/internal/primitive/enum"
	"order-service/internal/repositories/datastore/orders"
	"time"

	utilKafka "github.com/CROWNIX/go-utils/broker/kafka"

	productServiceDto "pkg/services/product-service/dto"

	"github.com/CROWNIX/go-utils/apperror"
	"github.com/CROWNIX/go-utils/databases/sqlx"
	"github.com/CROWNIX/go-utils/utils/generic"
	"github.com/segmentio/kafka-go"
)

func (s *orderService) CreateOrder(ctx context.Context, input CreateOrderServiceInput) (output *CreateOrderServiceOutput, err error) {
	productIDs := extractProductIDs(input.Items)

	err = s.tx.DoTxContext(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted}, func(ctx context.Context, tx sqlx.RDBMS) error {
		products, err := s.productService.GetProductByIds(ctx, productIDs)
		if err != nil {
			slog.Error("failed to get products", slog.String("error", err.Error()))
			return apperror.InternalServer("Internal Server Error")
		}

		if len(products) != len(productIDs) {
			slog.Error("one or more products not found", slog.Int("expected", len(productIDs)), slog.Int("found", len(products)))
			return apperror.NotFound("One of product not found")
		}

		err = validateOrderItems(input.Items, products)
		if err != nil{
			slog.Error("invalid order items", slog.String("error", err.Error()))
			return err
		}

		subTotal := calculateSubTotal(input.Items, products)

		if input.GrandTotal != subTotal {
			slog.Error("grand total mismatch", slog.Uint64("expected", subTotal), slog.Uint64("actual", input.GrandTotal))
			return apperror.BadRequest("Grand total does not match calculated total")
		}

		orderID, err := s.orderRepositoryWriter.CreateOrder(ctx, tx, orders.CreateOrderInput{
			UserID:         input.UserID,
			AddressID:      input.AddressID,
			PaymentMethodID: input.PaymentMethodID,
			Status:         string(enum.OrderStatusPendingPayment),
			GrandTotal:     subTotal,
		})
		
		if err != nil {
			slog.Error("failed to create order", slog.String("error", err.Error()))
			return apperror.InternalServer("Internal Server Error")
		}

		err = s.orderRepositoryWriter.CreateOrderDetail(ctx, orders.CreateOrderDetailInput{
			OrderID: orderID,
			Items:   generic.TransformSlice(input.Items, func(item Item) orders.Item {
				return orders.Item{
					ProductId: item.ProductId,
					Quantity:  item.Quantity,
				}
			}),
		})

		if err != nil {
			slog.Error("failed to create order detail", slog.String("error", err.Error()))
			return apperror.InternalServer("Internal Server Error")
		}

		event := map[string]any{
			"event_type": "order_created",
			"source": config.GetConfig().AppName,
			"timestamp": time.Now(),
			"data": map[string]any{
				"order_id": orderID,
			},
		}

		value, err := json.Marshal(event)
		if err != nil {
			slog.Error("failed to marshal event", slog.String("error", err.Error()))
			return apperror.InternalServer("Internal Server Error")
		}


		if s.pubSubKafka != nil {
			_, err = s.pubSubKafka.Publish(ctx, utilKafka.PubInput{
				KeyWriter: config.GetConfig().ServiceName,
				Messages: []kafka.Message{
					{
						Value: []byte(value),
						Topic: string(enum.KafkaTopicOrderCreated),
					},
				},
			})
			if err != nil {
				slog.Error("failed to publish event", slog.String("error", err.Error()))
				return apperror.InternalServer("Internal Server Error")
			}
		} else {
			slog.Warn("pubsub kafka is nil, skipping publish")
		}

		return nil
	})

	if err != nil {
		slog.Error("transaction failed", slog.String("error", err.Error()))
		return nil, err
	}

	return &CreateOrderServiceOutput{}, nil
}

func extractProductIDs(items []Item) []uint64 {
	return generic.TransformSlice(items, func(item Item) uint64 {
		return item.ProductId
	})
}

func validateOrderItems(items []Item, products []productServiceDto.GetDetailProductOutput) error {
	productOrderMap := map[uint64]Item{}
	for _, item := range items {
		productOrderMap[item.ProductId] = item
	}

	for _, product := range products {
		orderItem := productOrderMap[product.ID]
		if orderItem.Quantity > product.Stock {
			return apperror.BadRequest(fmt.Sprintf("Product '%s' is out of stock", product.Name))
		}

		if product.MaximumPurchase != nil {
			if orderItem.Quantity < uint32(product.MinimumPurchase) || orderItem.Quantity > uint32(*product.MaximumPurchase) {
				return apperror.BadRequest(fmt.Sprintf("Invalid quantity for product '%s'", product.Name))
			}
		} else if orderItem.Quantity < uint32(product.MinimumPurchase) {
			return apperror.BadRequest(fmt.Sprintf("Quantity below minimum purchase for product '%s'", product.Name))
		}
	}

	return nil
}

func calculateSubTotal(items []Item, products []productServiceDto.GetDetailProductOutput) uint64 {
	productMap := map[uint64]productServiceDto.GetDetailProductOutput{}
	for _, product := range products {
		productMap[product.ID] = product
	}

	var subTotal uint64
	for _, item := range items {
		product := productMap[item.ProductId]
		subTotal += uint64(item.Quantity) * product.Price
	}

	return subTotal
}


