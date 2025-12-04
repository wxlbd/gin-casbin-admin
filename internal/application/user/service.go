package user

import (
	"context"
	"time"

	"github.com/wxlbd/gin-casbin-admin/internal/domain/role"
	"github.com/wxlbd/gin-casbin-admin/internal/domain/user"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
	"github.com/wxlbd/gin-casbin-admin/pkg/jwtx"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Create(ctx context.Context, req *CreateUserRequest, createdBy uint64) error
	Update(ctx context.Context, req *UpdateUserRequest) error
	Delete(ctx context.Context, ids ...uint64) error
	FindByID(ctx context.Context, id uint64) (*UserResponse, error)
	FindByUsername(ctx context.Context, username string) (*UserResponse, error)
	List(ctx context.Context, req *UserListRequest) (*UserListResponse, error)
	UpdatePassword(ctx context.Context, id uint64, oldPassword, newPassword string) error
	ResetPassword(ctx context.Context, id uint64, newPassword string) error
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	RefreshToken(ctx context.Context, req *RefreshTokenRequest) (*LoginResponse, error)
	Logout(ctx context.Context, token string) error
	GetUserRoles(ctx context.Context, userID uint64) ([]*role.Role, error)
	AssignRoles(ctx context.Context, userID uint64, roleIDs []uint64) error
}

type userService struct {
	repo   user.Repository
	jwt    *jwtx.JWT
	logger *log.Logger
}

func NewUserService(logger *log.Logger, repo user.Repository, jwt *jwtx.JWT) Service {
	return &userService{
		repo:   repo,
		jwt:    jwt,
		logger: logger,
	}
}

func (s *userService) Create(ctx context.Context, req *CreateUserRequest, createdBy uint64) error {
	// 检查用户名是否存在
	existUser, err := s.repo.FindByUsername(ctx, req.Username)
	if err != nil && err != errors.ErrNotFound {
		return err
	}
	if existUser != nil {
		return errors.WithMsg(errors.AlreadyExists, "用户名已存在")
	}

	u := req.ToEntity(createdBy)
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.WithMsg(errors.ServerError, "密码加密失败")
	}
	u.Password = string(hashedPassword)

	return s.repo.Create(ctx, u)
}

func (s *userService) Update(ctx context.Context, req *UpdateUserRequest) error {
	existUser, err := s.repo.FindByID(ctx, req.ID)
	if err != nil {
		if err == errors.ErrNotFound {
			return errors.WithMsg(errors.NotFound, "用户不存在")
		}
		s.logger.Error("查询用户失败", zap.Error(err))
		return errors.ErrDatabase
	}

	// 如果修改了用户名，需要检查新用户名是否已存在
	if req.Username != existUser.Username {
		if exist, _ := s.repo.FindByUsername(ctx, req.Username); exist != nil {
			return errors.WithMsg(errors.AlreadyExists, "用户名已存在")
		}
	}

	u := req.ToEntity()
	// 保持原有密码如果不修改
	if u.Password == "" {
		u.Password = existUser.Password
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.WithMsg(errors.ServerError, "密码加密失败")
		}
		u.Password = string(hashedPassword)
	}

	// 保持其他未修改字段
	u.CreatedBy = existUser.CreatedBy
	u.CreatedAt = existUser.CreatedAt

	return s.repo.Update(ctx, u)
}

func (s *userService) Delete(ctx context.Context, ids ...uint64) error {
	return s.repo.Delete(ctx, ids...)
}

func (s *userService) FindByID(ctx context.Context, id uint64) (*UserResponse, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToUserResponse(u), nil
}

func (s *userService) FindByUsername(ctx context.Context, username string) (*UserResponse, error) {
	u, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return ToUserResponse(u), nil
}

func (s *userService) List(ctx context.Context, req *UserListRequest) (*UserListResponse, error) {
	query := req.ToQuery()
	users, total, err := s.repo.List(ctx, query)
	if err != nil {
		return nil, err
	}
	return ToUserListResponse(users, total), nil
}

func (s *userService) UpdatePassword(ctx context.Context, id uint64, oldPassword, newPassword string) error {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if err == errors.ErrNotFound {
			return errors.WithMsg(errors.NotFound, "用户不存在")
		}
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(oldPassword)); err != nil {
		return errors.WithMsg(errors.Unauthorized, "旧密码错误")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.WithMsg(errors.ServerError, "密码加密失败")
	}

	u.Password = string(hashedPassword)
	return s.repo.Update(ctx, u)
}

func (s *userService) ResetPassword(ctx context.Context, id uint64, newPassword string) error {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if err == errors.ErrNotFound {
			return errors.WithMsg(errors.NotFound, "用户不存在")
		}
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.WithMsg(errors.ServerError, "密码加密失败")
	}
	u.Password = string(hashedPassword)
	return s.repo.Update(ctx, u)
}

func (s *userService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	// TODO: Verify Captcha

	u, err := s.repo.FindByUsername(ctx, req.Username)
	if err != nil {
		if err == errors.ErrNotFound {
			return nil, errors.WithMsg(errors.NotFound, "用户不存在")
		}
		s.logger.Error("查询用户失败", zap.Error(err))
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		s.logger.Warn("密码错误", zap.Error(err))
		return nil, errors.WithMsg(errors.Unauthorized, "密码错误")
	}

	accessToken, refreshToken, err := s.jwt.GenerateToken(u.ID, u.Username)
	if err != nil {
		return nil, err
	}

	u.LoginTime = time.Now()
	// TODO: Update Login IP
	if err := s.repo.Update(ctx, u); err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expires:      "TODO", // Calculate expiration
	}, nil
}

func (s *userService) RefreshToken(ctx context.Context, req *RefreshTokenRequest) (*LoginResponse, error) {
	accessToken, refreshToken, err := s.jwt.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *userService) Logout(ctx context.Context, token string) error {
	claims, err := s.jwt.ParseToken(ctx, token, false)
	if err != nil {
		return err
	}
	return s.jwt.AddToBlacklist(ctx, token, claims)
}

func (s *userService) GetUserRoles(ctx context.Context, userID uint64) ([]*role.Role, error) {
	return s.repo.GetUserRoles(ctx, userID)
}

func (s *userService) AssignRoles(ctx context.Context, userID uint64, roleIDs []uint64) error {
	return s.repo.AssignRoles(ctx, userID, roleIDs)
}
