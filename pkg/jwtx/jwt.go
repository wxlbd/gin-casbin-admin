package jwtx

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/wxlbd/gin-casbin-admin/pkg/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type Claims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWT struct {
	config *config.JWTConfig
	redis  *redis.Client
	// 添加互斥锁，用于并发控制
	renewLock sync.Mutex
}

func New(cfg *config.Config, redis *redis.Client) *JWT {
	return &JWT{
		config: &cfg.JWT,
		redis:  redis,
	}
}

// GenerateToken 生成访问令牌（AccessToken）和刷新令牌（RefreshToken）。
// 该方法根据用户ID和用户名创建两个JWT令牌，每个令牌都有各自的过期时间和密钥。
// 参数:
//   - userID: 用户ID，用于标识令牌的拥有者。
//   - username: 用户名，用于在令牌中标识用户。
//
// 返
// 返回值:
//   - accessToken: 生成的访问令牌，用于用户身份验证。
//   - refreshToken: 生成的刷新令牌，用于获取新的访问令牌。
//   - err: 可能发生的错误，如果生成令牌失败。
func (j *JWT) GenerateToken(userID uint64, username string) (accessToken, refreshToken string, err error) {
	// 生成 Access Token
	accessClaims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.AccessExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.config.Issuer,
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(j.config.AccessSecret))
	if err != nil {
		return "", "", err
	}

	// 生成 Refresh Token
	refreshClaims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.RefreshExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.config.Issuer,
		},
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(j.config.RefreshSecret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// ParseToken 解析JWT令牌并验证其有效性。
// 该方法首先检查令牌是否在黑名单中（对于非刷新令牌而言），
// 然后使用相应的密钥解析令牌，最后验证令牌的有效性。
//
//		ctx context.Context: 上下文对象，用于传递请求范围的配置、超时设置等。
//		tokenString string: 待解析的JWT令牌字符串。
//		isRefreshToken bool: 指示令牌是否为刷新令牌的布尔值。
//
//	  isRefreshToken bool: 指示令牌是否为刷新令牌的布尔值。
//
//		*Claims: 如果令牌有效，返回一个包含令牌声明的指针。
//		error: 如果解析过程中发生错误或令牌无效，返回一个错误。
//	  error: 如果解析过程中发生错误或令牌无效，返回一个错误。
func (j *JWT) ParseToken(ctx context.Context, tokenString string, isRefreshToken bool) (*Claims, error) {
	// 检查是否在黑名单中
	if !isRefreshToken {
		inBlacklist, err := j.IsInBlacklist(ctx, tokenString)
		if err != nil {
			return nil, err
		}
		if inBlacklist {
			return nil, errors.New("token is in blacklist")
		}
	}

	// 根据令牌类型选择相应的密钥
	secret := j.config.AccessSecret
	if isRefreshToken {
		secret = j.config.RefreshSecret
	}

	// 使用选择的密钥解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	// 验证令牌的有效性
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	// 如果令牌无效，返回错误
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新访问令牌和刷新令牌。
// 该方法通过验证现有的刷新令牌来生成新的访问令牌和刷新令牌。
//
//		ctx - 上下文，用于传递请求范围的 deadline、取消信号等。
//		refreshToken - 需要刷新的刷新令牌字符串。
//
//	  refreshToken - 需要刷新的刷新令牌字符串。
//
//		两个字符串分别代表新生成的访问令牌和刷新令牌，以及一个错误对象，如果过程中发生错误则返回该错误。
//	  两个字符串分别代表新生成的访问令牌和刷新令牌，以及一个错误对象，如果过程中发生错误则返回该错误。
func (j *JWT) RefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	// 解析并验证刷新令牌，true 表示该令牌是刷新令牌。
	claims, err := j.ParseToken(ctx, refreshToken, true)
	if err != nil {
		// 如果解析或验证失败，返回空字符串和错误。
		return "", "", err
	}

	// 生成并返回新的访问令牌和刷新令牌。
	return j.GenerateToken(claims.UserID, claims.Username)
}

// 生成黑名单的 key
func (j *JWT) getBlacklistKey(tokenStr string) string {
	return "token:blacklist:" + tokenStr
}

// 生成续期记录的 key
func (j *JWT) getRenewalKey(userID uint64) string {
	return "token:renewal:" + fmt.Sprintf("%d", userID)
}

// AddToBlacklist 将指定的令牌添加到黑名单中。
// 此函数旨在防止已使用的令牌再次被验证和接受。
//
//		ctx - 上下文，用于传递请求范围的 deadline、取消信号等。
//		tokenStr - 需要添加到黑名单的令牌字符串。
//		claims - 包含令牌声明的结构体指针，用于获取令牌的过期时间和用户ID。
//
//	  claims - 包含令牌声明的结构体指针，用于获取令牌的过期时间和用户ID。
//
//		如果操作成功，则返回 nil；否则返回错误。
//	  如果操作成功，则返回 nil；否则返回错误。
func (j *JWT) AddToBlacklist(ctx context.Context, tokenStr string, claims *Claims) error {
	// 计算令牌的过期时间与当前时间的差值。
	expiration := time.Until(claims.ExpiresAt.Time)
	// 如果令牌已经过期，则记录日志并无需进一步操作。
	if expiration < 0 {
		log.Printf("token already expired for user %d", claims.UserID)
		return nil
	}

	// 使用 Redis 将令牌添加到黑名单中，设置过期时间为令牌的剩余有效期。
	if err := j.redis.Set(ctx, j.getBlacklistKey(tokenStr), "1", expiration).Err(); err != nil {
		// 如果添加失败，记录错误并返回。
		log.Printf("failed to add token to blacklist for user %d: %v", claims.UserID, err)
		return err
	}

	// 如果添加成功，记录日志。
	log.Printf("token added to blacklist for user %d", claims.UserID)
	return nil
}

// IsInBlacklist 检查给定的token是否在黑名单中。
// 该方法使用Redis来存储黑名单中的token，以实现高效的查询。
//
//		ctx context.Context: 上下文，用于传递请求范围的 deadline、取消信号等。
//		tokenStr string: 需要检查的token字符串。
//
//	  tokenStr string: 需要检查的token字符串。
//
//		bool: 如果token在黑名单中，则返回true；否则返回false。
//		error: 如果在检查过程中发生错误，则返回该错误。
//	  error: 如果在检查过程中发生错误，则返回该错误。
func (j *JWT) IsInBlacklist(ctx context.Context, tokenStr string) (bool, error) {
	// 使用Redis的Exists命令检查黑名单键是否存在。
	i, err := j.redis.Exists(ctx, j.getBlacklistKey(tokenStr)).Result()
	// 如果存在，i将大于0，表示token在黑名单中；否则，i为0。
	return i != 0, err
}

// CheckAndRenewToken 检查访问令牌的有效性，并在需要时续发新的令牌。
// 该方法主要解决了在高并发场景下，如何安全地续发令牌，同时防止令牌被重复续发的问题。
//
//		ctx - 上下文，用于传递请求范围的上下文信息。
//		tokenStr - 当前的访问令牌字符串。
//		claims - 包含令牌声明的指针，用于获取令牌的过期时间和用户信息。
//
//	  claims - 包含令牌声明的指针，用于获取令牌的过期时间和用户信息。
//
//		newAccessToken - 新生成的访问令牌，如果不需要续发，则为空字符串。
//		needRenew - 布尔值，指示是否需要续发令牌。
//		err - 错误对象，如果在检查或续发令牌过程中遇到错误，则返回该错误。
//	  err - 错误对象，如果在检查或续发令牌过程中遇到错误，则返回该错误。
func (j *JWT) CheckAndRenewToken(ctx context.Context, tokenStr string, claims *Claims) (newAccessToken string, needRenew bool, err error) {
	// 使用互斥锁确保并发安全
	j.renewLock.Lock()
	defer j.renewLock.Unlock()

	// 计算令牌剩余的有效时间
	remainingTime := time.Until(claims.ExpiresAt.Time)

	// 如果令牌剩余有效时间小于访问令牌总有效期的四分之一，则考虑续发新令牌
	if remainingTime < j.config.AccessExpire/4 {
		// 获取续发令牌的键
		renewalKey := j.getRenewalKey(claims.UserID)
		// 检查当前令牌是否已经有过续发记录
		exists, err := j.redis.Exists(ctx, renewalKey).Result()
		if err != nil {
			return "", false, fmt.Errorf("check renewal status failed: %w", err)
		}

		// 如果当前令牌没有续发记录，则生成新的访问令牌
		if exists != 0 {
			newAccessToken, err = j.generateAccessToken(claims.UserID, claims.Username)
			if err != nil {
				return "", false, fmt.Errorf("generate new token failed: %w", err)
			}

			// 使用 lua 脚本确保原子性
			script := `
                if redis.call("exists", KEYS[1]) == 0 then
                    redis.call("setex", KEYS[1], ARGV[1], "1")
                    return 1
                end
                return 0
            `

			// 执行 lua 脚本，尝试设置续发状态，并判断续发是否成功
			success, err := j.redis.Eval(ctx, script, []string{renewalKey}, int(remainingTime.Seconds())).Int64()
			if err != nil {
				return "", false, fmt.Errorf("set renewal status failed: %w", err)
			}

			// 如果续发成功，则返回新的访问令牌和续发状态
			if success == 1 {
				return newAccessToken, true, nil
			}
		}
	}

	// 如果不需要续发令牌，则返回空字符串和 false
	return "", false, nil
}

// generateAccessToken 生成访问令牌（AccessToken）。
// 该方法根据用户ID和用户名创建JWT令牌，包含令牌过期时间、签发时间和签发者等信息。
//
//		userID - 用户ID，用于标识令牌的拥有者。
//		username - 用户名，用于在令牌中标识用户。
//
//	  username - 用户名，用于在令牌中标识用户。
//
//		生成的JWT令牌字符串和可能发生的错误。
//	  生成的JWT令牌字符串和可能发生的错误。
func (j *JWT) generateAccessToken(userID uint64, username string) (string, error) {
	// 创建Claims结构体，包含用户ID、用户名和令牌的注册声明。
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置令牌过期时间为当前时间加上配置的访问令牌过期时长。
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.AccessExpire)),
			// 设置令牌签发时间为当前时间。
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 设置令牌的签发者为配置中的Issuer。
			Issuer: j.config.Issuer,
		},
	}

	// 使用HS256算法创建并签发JWT令牌，并返回签名后的令牌字符串。
	// 如果签发过程中出现错误，也会返回相应的错误。
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.config.AccessSecret))
}
