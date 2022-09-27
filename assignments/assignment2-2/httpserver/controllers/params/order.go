package params

type OrderCreateRequest struct {
	OrderCode    string `json:"order_code" validate:"required"`
	CustomerName string `json:"customer_name" validate:"required"`
	OrderedAt    string `json:"ordered_at" validate:"required"`
}

type OrderUpdateRequest struct {
	OrderCode    string `json:"order_code" validate:"required"`
	CustomerName string `json:"customer_name" validate:"required"`
	OrderedAt    string `json:"ordered_at" validate:"required"`
}
