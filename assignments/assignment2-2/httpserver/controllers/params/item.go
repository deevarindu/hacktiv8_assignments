package params

type ItemCreateRequest struct {
	ItemCode    string `json:"item_code" validate:"required"`
	Description string `json:"description" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
	OrderId     int    `json:"order_id" validate:"required"`
}

type ItemUpdateRequest struct {
	ItemCode    string `json:"item_code" validate:"required"`
	Description string `json:"description" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
	OrderId     int    `json:"order_id" validate:"required"`
}