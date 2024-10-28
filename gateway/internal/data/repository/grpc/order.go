package grpc

import (
	"fmt"
	"strconv"

	"github.com/superdev/ecommerce/gateway/internal/adapters/grpc/proto"
	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
)

type GrpcOrder struct {
	ID           uint64
	CustomerId   *uint32
	CustomerName *string
	Total        *uint64
	Items        []*GrpcOrderItem
}

func (us *GrpcOrder) GetId() string {
	return fmt.Sprintf("%d", us.ID)
}

func (us *GrpcOrder) GetCustomerId() *uint32 {
	return us.CustomerId
}

func (us *GrpcOrder) GetCustomerName() *string {
	return us.CustomerName
}

func (us *GrpcOrder) GetTotal() *uint64 {
	return us.Total
}

func (us *GrpcOrder) GetItems() []entity.OrderItem {
	var items []entity.OrderItem
	for _, pItem := range us.Items {
		items = append(items, pItem)
	}
	return items
}

//

type GrpcOrderItem struct {
	ProductId   uint64
	Price       int64
	Quantity    int32
	Name        *string
	Description *string
}

func (us *GrpcOrderItem) GetProductId() string {
	return fmt.Sprintf("%d", us.ProductId)
}
func (us *GrpcOrderItem) GetPrice() int64 {
	return us.Price
}

func (us *GrpcOrderItem) GetQuantity() int32 {
	return us.Quantity
}

//

func fromProtoToOrderEntity(pOrder *proto.Order) entity.Order {
	var order GrpcOrder
	order.ID = uint64(pOrder.Id)
	order.CustomerName = pOrder.CustomerName
	orderTotal := uint64(pOrder.TotalPrice)
	order.Total = &orderTotal
	if nil != pOrder.CustomerId {
		cID := uint32(*pOrder.CustomerId)
		order.CustomerId = &cID
	}
	var items []*GrpcOrderItem
	for _, pItem := range pOrder.Items {
		var item GrpcOrderItem
		item.ProductId = uint64(pItem.ProductId)
		item.Quantity = pItem.Quantity
		item.Name = pItem.Name
		item.Description = pItem.Description
		items = append(items, &item)
	}
	order.Items = items
	return &order
}

func fromOrderEntity(order entity.Order) *GrpcOrder {
	if gOrder, ok := order.(*GrpcOrder); ok {
		return gOrder
	} else {
		gOrder = &GrpcOrder{}
		ID, err := strconv.ParseUint(order.GetId(), 10, 64)
		if nil != err {
			// panic(err)
		} else {
			gOrder.ID = ID
		}

		return gOrder
	}
}

func toOrderEntity(order *GrpcOrder) entity.Order {
	return order
}
