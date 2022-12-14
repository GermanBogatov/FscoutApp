package service

import (
	"context"
	"github.com/GermanBogatov/auth_service/internal/model"
	"github.com/GermanBogatov/auth_service/internal/model/modelScout"
	"github.com/GermanBogatov/auth_service/internal/model/modelSportsman"
	"github.com/GermanBogatov/auth_service/internal/service/serviceAdmin"
	"github.com/GermanBogatov/auth_service/internal/service/serviceScout"
	"github.com/GermanBogatov/auth_service/internal/service/serviceSportsman"
	"github.com/GermanBogatov/auth_service/internal/storage"
	"github.com/GermanBogatov/auth_service/pkg/logging"
)

type AuthorizationSportsman interface {
	CreateSportsman(ctx context.Context, sportsman modelSportsman.SportsmanDTO) (string, error)
	SignInSportsman(ctx context.Context, sportsman model.SignInDTO) (model.AuthDTO, error)
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, token string) (int, error)
}

type AuthorizationScout interface {
	CreateScout(ctx context.Context, scout modelScout.ScoutDTO) (string, error)
	SignInScout(ctx context.Context, scout model.SignInDTO) (model.AuthDTO, error)
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, token string) (int, error)
}

type AuthorizationAdmin interface {
	CreateAdmin(ctx context.Context) (int, error)
	GetAdmin(ctx context.Context, username, password string) error
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, token string) (int, error)
}

type Service struct {
	AuthorizationSportsman
	AuthorizationScout
	AuthorizationAdmin
}

func NewService(storage *storage.Storage, logger logging.Logger) (*Service, error) {
	return &Service{
		AuthorizationSportsman: serviceSportsman.NewAuthServiceSportsman(storage.AuthorizationSportsman, logger),
		AuthorizationScout:     serviceScout.NewAuthServiceScout(storage.AuthorizationScout, logger),
		AuthorizationAdmin:     serviceAdmin.NewAuthServiceAdmin(storage.AuthorizationAdmin, logger),
	}, nil
}
