package admin

import (
	"ez/dto/request"
	"ez/dto/response"
	"ez/pkg/server"
	"ez/service"

	"github.com/gin-gonic/gin"
)

var userService = service.NewUserService()

// @Tags ADMIN User
// @Summary 创建用户
// @Description 创建用户
// @Accept json
// @Produce json
// @Param body body request.UserAdminCreateRequest true "请求参数"
// @Success 200 {object} response.CommonResponse
// @Router /admin/user/create [post]
func UserAdminCreate(ctx *gin.Context) {
	var err error

	var params request.UserAdminCreateRequest

	err = ctx.ShouldBindJSON(&params)

	if server.GinError(ctx, err, "参数缺失", 400) {
		return
	}

	err = userService.AdminCreate(params)

	if server.GinError(ctx, err, "创建失败", 500) {
		return
	}

	server.GinReply(ctx, "创建成功", 200, nil)
}

// @Tags ADMIN User
// @Summary 删除用户
// @Description 删除用户
// @Accept json
// @Produce json
// @Param body body request.UserAdminRemoveRequest true "请求参数"
// @Success 200 {object} response.CommonResponse
// @Router /admin/user/remove [post]
func UserAdminRemove(ctx *gin.Context) {
	var err error

	var params request.UserAdminRemoveRequest

	err = ctx.ShouldBindJSON(&params)

	if server.GinError(ctx, err, "参数缺失", 400) {
		return
	}

	err = userService.AdminRemove(params)

	if server.GinError(ctx, err, "删除失败", 500) {
		return
	}

	server.GinReply(ctx, "删除成功", 200, nil)
}

// @Tags ADMIN User
// @Summary 更新用户
// @Description 更新用户
// @Accept json
// @Produce json
// @Param body body request.UserAdminEditRequest true "请求参数"
// @Success 200 {object} response.CommonResponse
// @Router /admin/user/edit [post]
func UserAdminEdit(ctx *gin.Context) {
	var err error

	var params request.UserAdminEditRequest

	err = ctx.ShouldBindJSON(&params)

	if server.GinError(ctx, err, "参数缺失", 400) {
		return
	}

	err = userService.AdminEdit(params)

	if server.GinError(ctx, err, "编辑失败", 500) {
		return
	}

	server.GinReply(ctx, "编辑成功", 200, nil)
}

// @Tags ADMIN User
// @Summary 查询用户
// @Description 查询用户
// @Accept json
// @Produce json
// @Param body body request.UserAdminFindPagingRequest true "请求参数"
// @Success 200 {object} response.CommonResponse{data=response.UserAdminFindPagingResponse}
// @Router /admin/user/find/paging [post]
func UserAdminFindPaging(ctx *gin.Context) {
	var err error

	var params request.UserAdminFindPagingRequest

	err = ctx.ShouldBindJSON(&params)

	if server.GinError(ctx, err, "参数缺失", 400) {
		return
	}

	var data response.UserAdminFindPagingResponse
	data, err = userService.AdminFindPaging(params)

	if server.GinError(ctx, err, "查询失败", 500) {
		return
	}

	server.GinReply(ctx, "查询成功", 200, data)
}
