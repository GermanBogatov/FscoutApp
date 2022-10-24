package serviceSportsman

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
type AuthServiceSportsman struct {
	storage storage.AuthorizationSportsman
	logger  logging.Logger
}

func NewAuthServiceSportsman(storage storage.AuthorizationSportsman, logger logging.Logger) *AuthServiceSportsman {
	return &AuthServiceSportsman{
		storage: storage,
		logger:  logger,
	}
}

func (s *AuthServiceSportsman) CreateSportsman(ctx context.Context) (int, error) {

	return s.storage.CreateSportsman(ctx)
}

func (s *AuthServiceSportsman) GetSportsman(ctx context.Context, username, password string) error {
	return s.storage.GetSportsman(ctx, username, generatePasswordHash(password))
}
func (s *AuthServiceSportsman) GenerateToken(ctx context.Context, username, password string) (string, error) {
	return "", nil
}

func (s *AuthServiceSportsman) ParseToken(ctx context.Context, accessToken string) (int, error) {
	return 0, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
