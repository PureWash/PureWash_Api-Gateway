package domain

type PrimaryKey struct {
	ID string `bson:"id"`
}

type GetListRequest struct {
	Page   int32  `bson:"page"`
	Limit  int32  `bson:"limit"`
	Search string `bson:"search"`
}

type Response struct {
	StatusCode  int         `json:"status_code"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

type UpdatePasswordRequest struct {
	Login       string `bson:"login"`
	OldPassword string `bson:"old_password"`
	NewPassword string `bson:"new_password"`
}
