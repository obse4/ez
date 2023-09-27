package response

type UserAdminFindPagingResponse struct {
	Data  []UserAdminFindPagingResponseData `json:"data"`
	Total int64                             `json:"total"`
}

type UserAdminFindPagingResponseData struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	UserType  uint   `json:"user_type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserApiV1LoginResponse struct {
	Token     string `json:"token"`
	UserId    uint   `json:"user_id"`
	Username  string `json:"username"`
	UserType  uint   `json:"user_type"`
	ExpiresAt int64  `json:"expires_at"`
}
