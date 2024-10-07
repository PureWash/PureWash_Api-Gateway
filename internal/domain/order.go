package domain

//	type GetOrderResp struct {
//		ID         string  `json:"id"`
//		Client     Client  `json:"client"`
//		Service    Service `json:"service"`
//		Area       float32 `json:"area"`
//		TotalPrice float64 `json:"total_price"`
//		Status     string  `json:"status"`
//		CreatedAt  string  `json:"createdAt"`
//	}
//
//	type GetAllOrdersReq struct {
//		Fullname string `json:"full_name"`
//		Status   string `json:"status"`
//		OnTime   string `json:"on_time"`
//		OffSet   int32  `json:"offset"`
//		Limit    int32  `json:"limit"`
//	}
//
//	type GetOrdersResp struct {
//		Order      *[]pb.Order `json:"orders"`
//		OffSet     int32        `json:"offset"`
//		Limit      int32        `json:"limit"`
//		TotalCount int32        `json:"total_count"`
//	}
type OrderObject struct {
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
type OrderParams struct {
	ServiceId  string  `json:"service_id"`
	Area       float32 `json:"area"`
	TotalPrice float64 `json:"total_price"`
}
type CreateOrdResp struct {
	ClientInfo Client  `json:"client_info"`
	ID         string  `json:"id"`
	Area       float32 `json:"area"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"total_price"`
}

type UpdateOrderReq struct {
	ID          string  `json:"-"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	PhoneNumber string  `json:"phone_number"`
	Area        float32 `json:"area"`
	TotalPrice  float64 `json:"total_price"`
	Status      string  `json:"status"`
}
type UpdateOrderResp struct {
	ID         string  `json:"id"`
	Area       float32 `json:"area"`
	TotalPrice float32 `json:"total_price"`
	Status     string  `json:"status"`
	UpdatedAt  string  `json:"updated_at"`
}

type GetOrderResp struct {
	ID           string   `json:"id"`
	ClientInfo   Client   `json:"client_info"`
	ServicesInfo Services `json:"services_info"`
	Area         float32  `json:"area"`
	TotalPrice   float32  `json:"total_price"`
	Status       string   `json:"status"`
}

type GetAllOrderReq struct {
	FullName string `json:"full_name"`
	Status   string `json:"status"`
	Ontime   string `json:"ontime"`
	Page     int32  `json:"page"`
	Limit    int32  `json:"limit"`
}

type Order struct {
	ID         string `json:"id"`
	ClientInfo Client `json:"client_info"`
	Status     string `json:"status"`
}
type GetOrdersResp struct {
	OrderInfo  []Order `json:"order_info"`
	Page       int32   `json:"page"`
	Limit      int32   `json:"limit"`
	TotalCount int32   `json:"total_count"`
}
