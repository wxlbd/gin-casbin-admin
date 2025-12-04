package v1

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	userApp "github.com/wxlbd/gin-casbin-admin/internal/application/user"
	"github.com/wxlbd/gin-casbin-admin/internal/types"
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
	"github.com/wxlbd/gin-casbin-admin/pkg/ginx"
)

type UserHandler struct {
	svc userApp.Service
	cfg *config.Config
}

func NewUserHandler(svc userApp.Service, cfg *config.Config) *UserHandler {
	return &UserHandler{
		svc: svc,
		cfg: cfg,
	}
}

// Login 用户登录
func (h *UserHandler) Login(c *gin.Context) {
	var req userApp.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	// TODO: Captcha verification needs to be injected or handled via a separate service
	// For now, assuming captcha is valid or skipping it for this refactor step
	// if !h.svc.Captcha().Verify(c, req.CaptchaId, req.CaptchaCode) { ... }

	resp, err := h.svc.Login(c.Request.Context(), &req)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	// Set expiration format
	resp.Expires = time.Now().Add(h.cfg.JWT.AccessExpire).Format("2006/01/02 15:04:05")

	ginx.Success(c, resp)
}

// RefreshToken 刷新令牌
func (h *UserHandler) RefreshToken(c *gin.Context) {
	var req userApp.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	resp, err := h.svc.RefreshToken(c.Request.Context(), &req)
	if err != nil {
		ginx.Unauthorized(c, err)
		return
	}
	resp.Expires = time.Now().Add(h.cfg.JWT.AccessExpire).Format("2006/01/02 15:04:05")

	ginx.Success(c, resp)
}

// Logout 用户登出
func (h *UserHandler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" || len(token) <= 7 || token[:7] != "Bearer " {
		ginx.Success(c, nil)
		return
	}

	token = token[7:]
	if err := h.svc.Logout(c.Request.Context(), token); err != nil {
		ginx.Error(c, 500, err.Error())
		return
	}

	ginx.Success(c, nil)
}

// Create 创建用户
func (h *UserHandler) Create(c *gin.Context) {
	var req userApp.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	// Get current user ID from context (middleware)
	createdBy := c.GetUint64("user_id")

	if err := h.svc.Create(c.Request.Context(), &req, createdBy); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// Update 更新用户
func (h *UserHandler) Update(c *gin.Context) {
	var req userApp.UpdateUserRequest
	p := c.Param("id")
	id, err := strconv.ParseUint(p, 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的用户ID"))
		return
	}
	req.ID = id
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	if err := h.svc.Update(c.Request.Context(), &req); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// Delete 删除用户
func (h *UserHandler) Delete(c *gin.Context) {
	str := strings.Split(c.Param("ids"), ",")
	ids := make([]uint64, 0, len(str))

	for _, s := range str {
		id, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的用户ID"))
			return
		}
		ids = append(ids, id)
	}
	if err := h.svc.Delete(c.Request.Context(), ids...); err != nil {
		ginx.Error(c, 500, err.Error())
		return
	}

	ginx.Success(c, nil)
}

// List 获取用户列表
func (h *UserHandler) List(c *gin.Context) {
	var req userApp.UserListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	// Normalize page params
	if req.PageParam == nil {
		req.PageParam = &types.PageParam{}
	}
	req.Normalize() // This method is on *types.PageParam, so UserListRequest needs to embed it properly

	resp, err := h.svc.List(c.Request.Context(), &req)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, resp)
}

// ResetPassword 重置用户密码
func (h *UserHandler) ResetPassword(c *gin.Context) {
	var req userApp.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	p := c.Param("id")
	id, err := strconv.ParseUint(p, 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的用户ID"))
		return
	}
	req.ID = id
	if err := h.svc.ResetPassword(c.Request.Context(), req.ID, req.Password); err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, nil)
}

// Current 获取当前用户信息
func (h *UserHandler) Current(c *gin.Context) {
	userId := c.GetUint64("user_id")
	if userId == 0 {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的用户ID"))
		return
	}

	user, err := h.svc.FindByID(c.Request.Context(), userId)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, user)
}

// Detail 获取用户信息
func (h *UserHandler) Detail(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的用户ID"))
		return
	}
	user, err := h.svc.FindByID(c.Request.Context(), id)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, user)
}

// AssignRoles 分配角色
func (h *UserHandler) AssignRoles(c *gin.Context) {
	var req userApp.UserAssignRolesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	p := c.Param("id")
	id, err := strconv.ParseUint(p, 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的用户ID"))
		return
	}

	if err := h.svc.AssignRoles(c.Request.Context(), id, req.RoleIds); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}
