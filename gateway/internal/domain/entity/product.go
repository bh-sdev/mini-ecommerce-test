package entity

type Product interface {
	GetId() string
	GetName() string
	GetDescription() string
	GetQuantityAvailable() int64
	GetPrice() int64
}

//

func NewEmptyProduct() *product {
	return &product{}
}

func NewProductFromEntity(
	entity Product,
) *product {
	return &product{
		id:                entity.GetId(),
		name:              entity.GetName(),
		description:       entity.GetDescription(),
		price:             entity.GetPrice(),
		quantityAvailable: entity.GetQuantityAvailable(),
	}
}

type product struct {
	id                string
	name              string
	description       string
	quantityAvailable int64
	price             int64
}

func (s *product) GetId() string {
	return s.id
}

func (s *product) GetName() string {
	return s.name
}

func (s *product) SetName(name string) {
	s.name = name
}

func (s *product) GetDescription() string {
	return s.description
}

func (s *product) SetDescription(description string) {
	s.description = description
}

func (s *product) GetQuantityAvailable() int64 {
	return int64(s.quantityAvailable)
}

func (s *product) SetQuantityAvailable(quantityAvailable int64) {
	s.quantityAvailable = quantityAvailable
}

func (s *product) GetPrice() int64 {
	return s.price
}

func (s *product) SetPrice(price int64) {
	s.price = price
}
