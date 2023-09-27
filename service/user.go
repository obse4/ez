package service

import (
	"ez/common"
	"ez/dto/request"
	"ez/dto/response"
	"ez/global"
	"ez/model"
	"ez/pkg/hash"
	"fmt"
	"time"
)

type UserService interface {
	AdminCreate(p request.UserAdminCreateRequest) error
	AdminRemove(p request.UserAdminRemoveRequest) error
	AdminEdit(p request.UserAdminEditRequest) error
	AdminFindPaging(p request.UserAdminFindPagingRequest) (data response.UserAdminFindPagingResponse, err error)

	ApiV1Login(p request.UserApiV1LoginRequest) (data response.UserApiV1LoginResponse, err error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) AdminCreate(p request.UserAdminCreateRequest) error {
	db := global.Config.MysqlDb.Db.Model(&model.User{})

	var repeatCount int64
	db.Where("username = ?", p.Username).Count(&repeatCount)

	// username已存在
	if repeatCount > 0 {
		return fmt.Errorf("账户名已存在")
	}

	return db.Create(map[string]interface{}{
		"username":   p.Username,
		"user_type":  p.UserType,
		"password":   hash.String2Hash(p.Password),
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}).Error
}

func (s *userService) AdminRemove(p request.UserAdminRemoveRequest) error {
	db := global.Config.MysqlDb.Db.Model(&model.User{})

	var existCount int64
	db.Where("id = ?", p.Id).Count(&existCount)

	if existCount == 0 {
		return fmt.Errorf("账号不存在或已删除")
	}

	return db.Delete("id = ?", p.Id).Error
}

func (s *userService) AdminEdit(p request.UserAdminEditRequest) error {
	db := global.Config.MysqlDb.Db.Debug().Model(&model.User{})

	var exitCount int64
	db.Where("id <> ?", p.Id).Where("username = ?", p.Username).Count(&exitCount)

	if exitCount > 0 {
		return fmt.Errorf("账户名已存在")
	}

	updateCtx := map[string]interface{}{
		"username":   p.Username,
		"user_type":  p.UserType,
		"password":   hash.String2Hash(p.Password),
		"updated_at": time.Now(),
	}

	if p.Password == "" {
		updateCtx = map[string]interface{}{
			"username":   p.Username,
			"user_type":  p.UserType,
			"updated_at": time.Now(),
		}
	}

	return global.Config.MysqlDb.Db.Debug().Model(&model.User{}).Where("id = ?", p.Id).Updates(updateCtx).Error
}

func (s *userService) AdminFindPaging(p request.UserAdminFindPagingRequest) (data response.UserAdminFindPagingResponse, err error) {
	db := global.Config.MysqlDb.Db.Model(&model.User{})

	if p.UserType != 0 {
		db = db.Where("user_type = ?", p.UserType)
	}
	if p.Username != "" {
		db = db.Where("username like ?", fmt.Sprintf("%%%s%%", p.Username))
	}

	db.Count(&data.Total)

	var list []model.User

	db.Scopes(Paginate(p.PageIndex, p.PageSize)).Order("id ASC").Find(&list)

	for _, v := range list {
		data.Data = append(data.Data, response.UserAdminFindPagingResponseData{
			Id:        v.Id,
			Username:  v.Username,
			UserType:  v.UserType,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return
}

func (s *userService) ApiV1Login(p request.UserApiV1LoginRequest) (data response.UserApiV1LoginResponse, err error) {
	var user model.User
	global.Config.MysqlDb.Db.Debug().Model(&model.User{}).Where("username = ?", p.Username).First(&user)

	if user.Id == 0 {
		err = fmt.Errorf("用户账户或密码错误")
		return
	}

	if !hash.CompareHash(user.Password, p.Password) {
		err = fmt.Errorf("用户账户或密码错误")
		return
	}

	// 生成jwt
	userJwt := common.UserJwt{
		Username: user.Username,
		UserId:   user.Id,
		UserType: user.UserType,
	}
	data.Token, err = global.Config.Jwt.Jwt.CreateJwtToken(&userJwt)

	if err != nil {
		return
	}
	data.UserId = user.Id
	data.UserType = user.UserType
	data.Username = user.Username
	data.ExpiresAt = userJwt.ExpiresAt

	return
}
