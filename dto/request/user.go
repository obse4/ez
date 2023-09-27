package request

type UserAdminCreateRequest struct {
	// 用户账户
	Username string `json:"username"`
	// 用户类型 1-admin，2-user
	UserType uint `json:"user_type"`
	// 用户密码
	Password string `json:"password"`
}

type UserAdminRemoveRequest struct {
	// 用户id
	Id uint `json:"id"`
}

type UserAdminEditRequest struct {
	// 用户id
	Id uint `json:"id"`
	// 用户账户
	Username string `json:"username"`
	// 用户类型 1-admin，2-user
	UserType uint `json:"user_type"`
	// 用户密码
	Password string `json:"password"`
}

type UserAdminFindPagingRequest struct {
	FindPagingCommonRequest
	// 用户账户
	Username string `json:"username"`
	// 用户类型 1-admin，2-user
	UserType uint `json:"user_type"`
}

type UserApiV1LoginRequest struct {
	// 用户账户
	Username string `json:"username"`
	// 密码
	Password string `json:"password"`
}
