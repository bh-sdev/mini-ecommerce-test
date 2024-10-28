package grpc

import (
	"context"
	"errors"
	"strconv"

	"github.com/superdev/ecommerce/gateway/internal/adapters/grpc/proto"
	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/internal/domain/repository"
	"go.uber.org/zap"
)

type orderRepository struct {
	client proto.OrderServiceClient
	logger zap.Logger
}

func NewGrpcOrderRepository(
	client proto.OrderServiceClient,
	logger zap.Logger,
) repository.OrderRepository {
	return &orderRepository{
		client: client,
		logger: logger,
	}
}

func (s *orderRepository) Create(ctx context.Context, entity entity.Order) (entity.Order, error) {
	var req proto.CreateOrderRequest
	for _, item := range entity.GetItems() {
		var itReq proto.OrderItem
		productId, err := strconv.ParseUint(item.GetProductId(), 10, 32)
		if err != nil {
			panic(err)
		}
		itReq.ProductId = int32(productId)
		itReq.Quantity = item.GetQuantity()
		req.Items = append(req.Items, &itReq)
	}

	if order, err := s.client.CreateOrder(context.Background(), &req); err != nil {
		return nil, err
	} else {
		return fromProtoToOrderEntity(order.Order), nil
	}
}

func (s *orderRepository) Delete(ctx context.Context, entity entity.Order) error {
	id, err := strconv.ParseUint(entity.GetId(), 10, 32)
	if err != nil {
		return err
	}
	var req proto.DeleteOrderRequest
	req.Id = int32(id)
	_, err = s.client.DeleteOrder(context.Background(), &req)
	return err
}
func (s *orderRepository) FindAll(ctx context.Context) ([]entity.Order, error) {
	var req proto.Empty
	res, err := s.client.ListOrders(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	var orders = make([]entity.Order, 0)
	for _, ord := range res.Orders {
		order := fromProtoToOrderEntity(ord)
		orders = append(orders, order)
	}
	return orders, err

}

func (s *orderRepository) FindById(ctx context.Context, id string) (entity.Order, error) {
	_id, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}
	var req = proto.OrderRequest{
		Id: int32(_id),
	}
	if res, err := s.client.GetOrder(ctx, &req); err != nil {
		return nil, err
	} else {
		product := fromProtoToOrderEntity(res)
		return product, nil
	}
}

func (s *orderRepository) Update(ctx context.Context, entity entity.Order) error {
	return errors.ErrUnsupported
}
