package grpc

import (
	"fmt"
	"strconv"

	"github.com/superdev/ecommerce/gateway/internal/adapters/grpc/proto"
	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
)

type GrpcProduct struct {
	ID                uint64
	Name              string
	Description       string
	Price             float32
	AvailableQuantity int32
}

func (us *GrpcProduct) GetId() string {
	return fmt.Sprintf("%d", us.ID)
}

func (us *GrpcProduct) GetName() string {
	return us.Name
}
func (us *GrpcProduct) GetDescription() string {
	return us.Description
}
func (us *GrpcProduct) GetQuantityAvailable() int64 {
	return int64(us.AvailableQuantity)
}
func (us *GrpcProduct) GetPrice() int64 {
	return int64(us.Price)
}

func fromProtoToProductEntity(pProduct *proto.Product) entity.Product {
	// fmt.Printf("FOUND AVAILABLE QUANTITY: %d\n", pProduct.AvailableQuantity)
	var product GrpcProduct
	product.ID = uint64(pProduct.Id)
	product.Name = pProduct.Name
	product.Description = pProduct.Description
	product.AvailableQuantity = pProduct.AvailableQuantity
	product.Price = pProduct.Price
	return &product
}

func fromProductEntity(product entity.Product) *GrpcProduct {
	if gProduct, ok := product.(*GrpcProduct); ok {
		return gProduct
	} else {
		gProduct = &GrpcProduct{}

		ID, err := strconv.ParseUint(product.GetId(), 10, 64)
		if nil != err {
			// panic(err)
		} else {
			gProduct.ID = ID
		}
		gProduct.ID = ID
		gProduct.Name = product.GetName()
		gProduct.Description = product.GetDescription()
		gProduct.AvailableQuantity = int32(product.GetQuantityAvailable())
		gProduct.Price = float32(product.GetPrice())

		return gProduct
	}
}

func toProductEntity(product *GrpcProduct) entity.Product {
	return product
}
