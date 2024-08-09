package service

import (
	"context"
	v1 "gin-casbin-admin/api/v1"
	"gin-casbin-admin/internal/model"
	"gin-casbin-admin/internal/repository"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

type UserService interface {
	Add(ctx context.Context, req *v1.AddAdminUserRequest) error
	Login(ctx context.Context, req *v1.LoginRequest) (*Token, error)
	GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error
	SetUserRoles(ctx *gin.Context, v *v1.SetUserRoleRequest) error
}

func NewUserService(
	service *Service,
	userRepo repository.UserRepository,
	tokenRepo repository.TokenRepository,
	enforcer *casbin.Enforcer,
	conf *viper.Viper,
) UserService {
	return &userService{
		userRepo:  userRepo,
		Service:   service,
		enforcer:  enforcer,
		tokenRepo: tokenRepo,
		conf:      conf,
	}
}

type userService struct {
	conf      *viper.Viper
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
	enforcer  *casbin.Enforcer
	*Service
}

func (s *userService) SetUserRoles(ctx *gin.Context, v *v1.SetUserRoleRequest) error {
	if _, err2 := s.enforcer.DeleteRolesForUser(v.UserId); err2 != nil {
		return err2
	}
	_, err := s.enforcer.AddRolesForUser(v.UserId, v.RoleIds)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) Add(ctx context.Context, req *v1.AddAdminUserRequest) error {
	// check username
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return v1.ErrInternalServerError
	}
	if user != nil {
		return v1.ErrUsernameAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// 生成用户ID
	userId, err := s.sid.GenString()
	if err != nil {
		return err
	}
	user = &model.User{
		UserId:   userId,
		Email:    req.Email,
		Username: req.Username,
		Password: string(hashedPassword),
	}
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		// Create a user
		if err = s.userRepo.Create(ctx, user); err != nil {
			return err
		}
		if len(req.RoleIds) != 0 {
			if _, err = s.enforcer.AddRolesForUser(user.UserId, req.RoleIds); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest) (*Token, error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil || user == nil {
		return nil, v1.ErrUnauthorized
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	token, err := s.jwt.GenToken(user.UserId, user.Username, time.Now().Add(s.conf.GetDuration("security.jwt.access_token_expire")))
	if err != nil {
		return nil, err
	}
	refreshExpire := s.conf.GetDuration("security.jwt.refresh_token_expire")
	refreshToken, err := s.jwt.GenToken(user.UserId, user.Username, time.Now().Add(refreshExpire))
	if err != nil {
		return nil, err
	}
	if err := s.tokenRepo.StoreRefreshToken(ctx, user.UserId, refreshToken, refreshExpire); err != nil {
		return nil, err
	}
	return &Token{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetProfileResponseData{
		UserId:   user.UserId,
		Username: user.Username,
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	user.Email = req.Email
	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.userRepo.GetByEmail(ctx, email)
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.userRepo.GetByUsername(ctx, username)
}
