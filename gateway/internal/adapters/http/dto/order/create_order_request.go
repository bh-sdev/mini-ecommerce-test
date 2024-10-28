package order

type CreateOrderRequest struct {
	Items []CreateOrderRequestItem `json:"items" binding:"required"`
}

type CreateOrderRequestItem struct {
	ProductId string `json:"productId" binding:"required"`
	Quantity  int32  `json:"quantity" binding:"required"`
}
