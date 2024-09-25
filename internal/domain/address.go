package domain

type Address struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updatedAt"`
}

type AddressRequest struct {
	UserID    string `json:"user_id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
type AddressForUserRequest struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type AddressForUser struct {
	ID        string `json:"id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updatedAt"`
}
