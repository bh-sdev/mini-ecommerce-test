package product

type UpdateProductRequest struct {
	Name              *string `json:"name" binding:""`
	Description       *string `json:"description" binding:""`
	Price             *uint64 `json:"price" binding:""`
	AvailableQuantity *int64  `json:"availableQuantity" binding:""`
}
