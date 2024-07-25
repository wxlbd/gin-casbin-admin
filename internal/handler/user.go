package handler

import (
	"gin-casbin-admin/api/v1"
	"gin-casbin-admin/internal/service"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type AdminUserHandler struct {
	*Handler
	userService service.UserService
	enforcer    *casbin.Enforcer
}

func NewUserHandler(handler *Handler, userService service.UserService, enforcer *casbin.Enforcer) *AdminUserHandler {
	return &AdminUserHandler{
		Handler:     handler,
		userService: userService,
		enforcer:    enforcer,
	}
}

// AddAdminUser godoc
// @Summary 添加管理员
// @Schemes
// @Description 添加管理员
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body v1.AddAdminUserRequest true "params"
// @Success 200 {object} v1.Response
// @Router /register [post]
func (h *AdminUserHandler) AddAdminUser(ctx *gin.Context) {
	req := new(v1.AddAdminUserRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userService.Add(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("userService.AddAdminUser error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// Login godoc
// @Summary 账号登录
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body v1.LoginRequest true "params"
// @Success 200 {object} v1.LoginResponse
// @Router /login [post]
func (h *AdminUserHandler) Login(ctx *gin.Context) {
	var req v1.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	token, err := h.userService.Login(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	v1.HandleSuccess(ctx, v1.LoginResponseData{
		AccessToken: token,
	})
}

// GetProfile godoc
// @Summary 获取用户信息
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} v1.GetProfileResponse
// @Router /user [get]
func (h *AdminUserHandler) GetProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	v1.HandleSuccess(ctx, user)
}

// UpdateProfile godoc
// @Summary 修改用户信息
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body v1.UpdateProfileRequest true "params"
// @Success 200 {object} v1.Response
// @Router /user [put]
func (h *AdminUserHandler) UpdateProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)

	var req v1.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

func (h *AdminUserHandler) SetUserRoles(ctx *gin.Context) {

	var req v1.SetUserRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	//v, ok := ctx.Get("claims")
	//if !ok {
	//	v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
	//	return
	//}
	//claims, ok := v.(*jwt.MyCustomClaims)
	//if !ok {
	//	v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
	//	return
	//}
	//req.UserId = claims.UserId
	if err := h.userService.SetUserRoles(ctx, &req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}
