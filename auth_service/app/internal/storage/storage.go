package storage

import (
	"context"
	"github.com/GermanBogatov/auth_service/internal/model"
	"github.com/GermanBogatov/auth_service/internal/model/modelSportsman"
	"github.com/GermanBogatov/auth_service/internal/storage/storageAdmin"
	"github.com/GermanBogatov/auth_service/internal/storage/storageScout"
	"github.com/GermanBogatov/auth_service/internal/storage/storageSportsman"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/GermanBogatov/auth_service/pkg/postgresql"
)

type AuthorizationSportsman interface {
	CreateSportsman(ctx context.Context, sportsman modelSportsman.SportsmanDTO) (string, error)
	GetSportsman(ctx context.Context, sportsmanSign modelSportsman.SignInDTO) (model.AuthDTO, error)
}

type AuthorizationScout interface {
	CreateScout(ctx context.Context) (int, error)
	GetScout(ctx context.Context, username, password string) error
}

type AuthorizationAdmin interface {
	CreateAdmin(ctx context.Context) (int, error)
	GetAdmin(ctx context.Context, username, password string) error
}

type Storage struct {
	AuthorizationSportsman
	AuthorizationScout
	AuthorizationAdmin
}

func NewStorage(client postgresql.ClientPostgres, logger logging.Logger) *Storage {
	return &Storage{
		AuthorizationSportsman: storageSportsman.NewAuthorization(client, logger),
		AuthorizationScout:     storageScout.NewAuthorization(client, logger),
		AuthorizationAdmin:     storageAdmin.NewAuthorization(client, logger),
	}
}
