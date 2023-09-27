package model

type User struct {
	CommonModel
	Id       uint   `json:"id" gorm:"primarykey"`
	Username string `json:"username" gorm:"index; type:varchar(255); not null; comment:用户账号"`
	UserType uint   `json:"user_type" gorm:"not null; default:2; comment:用户类型1-admin,2-user"`
	Password string `json:"password" gorm:"type:varchar(500); not null; comment:用户密码"`
}
