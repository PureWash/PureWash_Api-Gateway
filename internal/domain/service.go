package domain

type Service struct {
	ID          string  `json:"id"`
	Traffic     string  `json:"traffic"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updated_at"`
}

type ServiceRequest struct {
	Traffic     string  `json:"traffic"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
type ServicesResponse struct {
	Services []*Service `json:"services"`
	Count    int64      `json:"count"`
	Page     int64      `json:"page"`
}
