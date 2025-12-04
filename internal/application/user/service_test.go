package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wxlbd/gin-casbin-admin/internal/domain/role"
	"github.com/wxlbd/gin-casbin-admin/internal/domain/user"
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
	"github.com/wxlbd/gin-casbin-admin/pkg/jwtx"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
)

// MockRepository is a mock implementation of user.Repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(ctx context.Context, user *user.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockRepository) Update(ctx context.Context, user *user.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockRepository) Delete(ctx context.Context, ids ...uint64) error {
	args := m.Called(ctx, ids)
	return args.Error(0)
}

func (m *MockRepository) FindByID(ctx context.Context, id uint64) (*user.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockRepository) List(ctx context.Context, query *user.UserQuery) ([]*user.User, int64, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]*user.User), args.Get(1).(int64), args.Error(2)
}

func (m *MockRepository) FindAll(ctx context.Context) ([]*user.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*user.User), args.Error(1)
}

func (m *MockRepository) GetUserRoles(ctx context.Context, userID uint64) ([]*role.Role, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*role.Role), args.Error(1)
}

func (m *MockRepository) AssignRoles(ctx context.Context, userID uint64, roleIDs []uint64) error {
	args := m.Called(ctx, userID, roleIDs)
	return args.Error(0)
}

func TestUserService_AssignRoles(t *testing.T) {
	mockRepo := new(MockRepository)
	logger := log.NewLog(&config.LogConfig{
		LogLevel: "debug",
		Encoding: "console",
	})
	jwt := &jwtx.JWT{} // Dummy JWT
	svc := NewUserService(logger, mockRepo, jwt)

	ctx := context.Background()
	userID := uint64(1)
	roleIDs := []uint64{1, 2, 3}

	// Expectation
	mockRepo.On("AssignRoles", ctx, userID, roleIDs).Return(nil)

	// Execute
	err := svc.AssignRoles(ctx, userID, roleIDs)

	// Verify
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
