package jwt

import (
	"context"
	"encoding/json"
	"github.com/GermanBogatov/auth_service/internal/config"
	"github.com/GermanBogatov/auth_service/internal/model"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

var _ Helper = &helper{}

type UserClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

type RT struct {
	RefreshToken string `json:"refresh_token"`
}

type helper struct {
	Logger      logging.Logger
	clientRedis *redis.Client
}

func NewHelper(logger logging.Logger, client *redis.Client) Helper {
	return &helper{
		Logger:      logger,
		clientRedis: client,
	}
}

type Helper interface {
	GenerateAccessToken(u model.AuthDTO) (string, string, error)
	UpdateRefreshToken(refreshToken string) (string, string, error)
}

func (h *helper) UpdateRefreshToken(refreshToken string) (string, string, error) {

	defer h.clientRedis.Del(context.Background(), refreshToken)

	userBytes := h.clientRedis.Get(context.Background(), refreshToken)

	var u model.AuthDTO
	err := json.Unmarshal([]byte(userBytes.Val()), &u)
	if err != nil {
		return "", "", err
	}

	return h.GenerateAccessToken(u)
	return "", "", nil

}

func (h *helper) GenerateAccessToken(u model.AuthDTO) (string, string, error) {
	key := []byte(config.GetConfig().JWT.Secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        u.Uuid,
			Audience:  u.Role,
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
		Email: u.Email,
	})

	accessToken, err := token.SignedString(key)
	if err != nil {
		return "", "", err
	}

	h.Logger.Info("create refresh token")
	refreshTokenUuid := uuid.New()
	userBytes, _ := json.Marshal(u)
	h.clientRedis.Set(context.Background(), refreshTokenUuid.String(), userBytes, 0)

	return accessToken, refreshTokenUuid.String(), err
}
