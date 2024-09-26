package domain

type Company struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CompanyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type CompaniesResponse struct {
	Companies []*Company `json:"companies"`
	Count     int64      `json:"count"`
	Page      int64      `json:"page"`
}
