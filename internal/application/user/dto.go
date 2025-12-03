package user

import (
	"github.com/wxlbd/gin-casbin-admin/internal/domain/user"
	"github.com/wxlbd/gin-casbin-admin/internal/types"
)

// UserBase 基础字段
type UserBase struct {
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Status   int8   `json:"status"`
	UserType int    `json:"user_type"`
	Signed   string `json:"signed"`
	Remark   string `json:"remark"`
}

// CreateUserRequest CreateUserReq 创建用户请求
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserBase
}

// UpdateUserRequest UpdateUserReq 更新用户请求
type UpdateUserRequest struct {
	ID       uint64 `json:"id" binding:"required"`
	Password string `json:"password,omitempty"` // 更新时密码可选
	UserBase
}

// UserResponse 用户响应
type UserResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	UserBase
	LoginIp   string `json:"login_ip"`
	LoginTime string `json:"login_time"`
	CreatedBy uint64 `json:"created_by"`
	UpdatedBy uint64 `json:"updated_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ToEntity 转换方法
func (req *CreateUserRequest) ToEntity(createdBy uint64) *user.User {
	return &user.User{
		Username:  req.Username,
		Password:  req.Password,
		Nickname:  req.Nickname,
		Phone:     req.Phone,
		Email:     req.Email,
		Status:    req.Status,
		UserType:  req.UserType,
		CreatedBy: createdBy,
	}
}

func (req *UpdateUserRequest) ToEntity() *user.User {
	return &user.User{
		ID:       req.ID,
		Password: req.Password, // 如果为空则不会更新
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
		Status:   req.Status,
		Username: req.Username,
		Remark:   req.Remark,
	}
}

// ToUserResponse 模型转响应
func ToUserResponse(m *user.User) *UserResponse {
	if m == nil {
		return nil
	}
	return &UserResponse{
		ID:       m.ID,
		Username: m.Username,
		UserBase: UserBase{
			Nickname: m.Nickname,
			Phone:    m.Phone,
			Email:    m.Email,
			Status:   m.Status,
			Remark:   m.Remark,
			Avatar:   m.Avatar,
			UserType: m.UserType,
			Signed:   m.Signed,
		},
		CreatedBy: m.CreatedBy,
		UpdatedBy: m.UpdatedBy,
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// ToUserResponseList 将用户列表转换为响应 DTO 列表
func ToUserResponseList(users []*user.User) []*UserResponse {
	if users == nil {
		return nil
	}
	list := make([]*UserResponse, 0, len(users))
	for _, u := range users {
		list = append(list, ToUserResponse(u))
	}
	return list
}

// ToUserListResponse 转换为分页响应
func ToUserListResponse(users []*user.User, total int64) *UserListResponse {
	return &UserListResponse{
		List:  ToUserResponseList(users),
		Total: total,
	}
}

// LoginRequest 用户登录请求
type LoginRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CaptchaId   string `json:"captcha_id" binding:"required"`   // 验证码ID
	CaptchaCode string `json:"captcha_code" binding:"required"` // 验证码值
}

// RefreshTokenRequest 刷新令牌请求
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// UserListRequest 用户列表请求
type UserListRequest struct {
	*types.PageParam
	Username string `form:"username"`
	Nickname string `form:"nickname"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Status   int8   `form:"status"`
}

func (req *UserListRequest) ToQuery() *user.UserQuery {
	return &user.UserQuery{
		Page:     req.Page,
		PageSize: req.PageSize,
		Username: req.Username,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
		Status:   req.Status,
	}
}

// ResetPasswordRequest 修改密码请求
type ResetPasswordRequest struct {
	ID       uint64
	Password string `json:"password" binding:"required,min=6"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Expires      string `json:"expires"`
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	List  []*UserResponse `json:"list"`
	Total int64           `json:"total"`
}

// UserAssignRolesRequest 用户分配角色
type UserAssignRolesRequest struct {
	RoleIds []uint64 `json:"roleIds"`
}
type UserRoleItem struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
type UserRolesResponse struct {
	Roles []*UserRoleItem `json:"roles"`
}
