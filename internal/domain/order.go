package domain

type Order struct {
	ID         string  `json:"id"`
	UserID     string  `json:"user_id"`
	ServiceID  string  `json:"service_id"`
	Area       float32 `json:"area"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
	CreatedAt  string  `json:"createdAt"`
}

type OrderRequest struct {
	UserID     string  `json:"user_id"`
	ServiceID  string  `json:"service_id"`
	Area       float32 `json:"area"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
}
type OrdersResponse struct {
	Orders []*Order `json:"orders"`
	Count  int64    `json:"count"`
	Page   int64    `json:"page"`
}
type OrderForUserRequest struct {
	ServiceID  string  `json:"service_id"`
	Area       float32 `json:"area"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
}
type OrderForUser struct {
	ID         string  `json:"id"`
	ServiceID  string  `json:"service_id"`
	Area       float32 `json:"area"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
	CreatedAt  string  `json:"createdAt"`
}
