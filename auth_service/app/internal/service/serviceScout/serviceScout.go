package serviceScout

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/GermanBogatov/auth_service/internal/storage"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	salt       = "sad342mslfd23412sdfsdf1234hgf"
	signingKey = ("HellowGerman! this is gin rest api")
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
type AuthServiceScout struct {
	storage storage.AuthorizationScout
	logger  logging.Logger
}

func NewAuthServiceScout(storage storage.AuthorizationScout, logger logging.Logger) *AuthServiceScout {
	return &AuthServiceScout{
		storage: storage,
		logger:  logger,
	}
}

func (s *AuthServiceScout) CreateScout(ctx context.Context) (int, error) {

	return s.storage.CreateScout(ctx)
}

func (s *AuthServiceScout) GetScout(ctx context.Context, username, password string) error {
	return s.storage.GetScout(ctx, username, generatePasswordHash(password))
}
func (s *AuthServiceScout) GenerateToken(ctx context.Context, username, password string) (string, error) {
	return "", nil
}

func (s *AuthServiceScout) ParseToken(ctx context.Context, accessToken string) (int, error) {
	return 0, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
