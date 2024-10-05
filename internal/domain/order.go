package domain

import pb "api_gateway/genproto/pure_wash"

type CreateOrderReq struct {
	UserID     string  `json:"user_id"`
	Client     Client  `json:"client"`
	ServiceID  string  `json:"service_id"`
	Area       float32 `json:"area"`
	TotalPrice float64 `json:"total_price"`
}

type CreateOrderResp struct {
	Fullname    string  `json:"full_name"`
	PhoneNumber string  `json:"phone_number"`
	Area        float32 `json:"area"`
	TotalPrice  float64 `json:"total_price"`
	CreatedAt   string  `json:"createdAt"`
}
type UpdateOrderReq struct {
	Orders []*Order `json:"orders"`
	Count  int64    `json:"count"`
	Page   int64    `json:"page"`
}
type UpdateOrderResp struct {
	Latitude   float32 `json:"latitude"`
	Longitude  float32 `json:"longitude"`
	Area       float32 `json:"area"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"total_price"`
}
type GetOrderResp struct {
	ID         string  `json:"id"`
	Client     Client  `json:"client"`
	Service    Service `json:"service"`
	Area       float32 `json:"area"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
	CreatedAt  string  `json:"createdAt"`
}
type GetAllOrdersReq struct {
	Fullname string `json:"full_name"`
	Status   string `json:"status"`
	OnTime   string `json:"on_time"`
	OffSet   int32  `json:"offset"`
	Limit    int32  `json:"limit"`
}

type GetOrdersResp struct {
	Order      *[]pb.Order `json:"orders"`
	OffSet     int32        `json:"offset"`
	Limit      int32        `json:"limit"`
	TotalCount int32        `json:"total_count"`
}
type Order struct {
	ID     string `json:"id"`
	Client Client `json:"client"`
	Status string `json:"status"`
}
type Client struct {
	Fullname    string  `json:"full_name"`
	PhoneNumber string  `json:"phone_number"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
}

type Services struct {
	Name    string `json:"name"`
	Tariffs string `json:"tariffs"`
}
