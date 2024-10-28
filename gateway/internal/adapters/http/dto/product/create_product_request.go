package product

type CreateProductRequest struct {
	Name              string `json:"name" binding:"required"`
	Description       string `json:"description" binding:"required"`
	Price             uint64 `json:"price" binding:"required"`
	AvailableQuantity int64  `json:"availableQuantity" binding:"required"`
}
