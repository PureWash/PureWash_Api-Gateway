package domain

type Order struct {
	ID         string  `json:"id"`
	UserID     string  `json:"user_id"`
	ServiceID  string  `json:"service_id"`
	Area       float32 `json:"area"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
	AddressID  string  `json:"address_id"`
	CreatedAt  string  `json:"createdAt"`
}

type OrderRequest struct {
	UserID     string  `json:"user_id"`
	ServiceID  string  `json:"service_id"`
	Area       float32 `json:"area"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
	AddressID  string  `json:"address_id"`
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
	AddressID  string  `json:"address_id"`
}
type OrderForUser struct {
	ID         string  `json:"id"`
	ServiceID  string  `json:"service_id"`
	Area       float32 `json:"area"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
	AddressID  string  `json:"address_id"`
	CreatedAt  string  `json:"createdAt"`
}
