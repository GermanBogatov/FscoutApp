package storageSportsman

import (
	"context"
	"fmt"
	"github.com/GermanBogatov/auth_service/internal/model/modelSportsman"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/GermanBogatov/auth_service/pkg/postgresql"
	"github.com/jackc/pgconn"
)

const (
	roleSportsman = "sportsman"
)

type repositoryAuthSportsman struct {
	client postgresql.ClientPostgres
	logger logging.Logger
}

func NewAuthorization(client postgresql.ClientPostgres, logger logging.Logger) *repositoryAuthSportsman {
	return &repositoryAuthSportsman{
		client: client,
		logger: logger,
	}
}

func (r *repositoryAuthSportsman) CreateSportsman(ctx context.Context, sportsman modelSportsman.SportsmanDTO) (string, error) {
	tx, err := r.client.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}
	defer tx.Commit(ctx)

	role := `
				SELECT role_uuid
				FROM role
				WHERE name=$1
			`
	if err := tx.QueryRow(ctx, role, roleSportsman).Scan(&sportsman.Role_uuid); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return "", newErr
		}
		return "", err
	}

	q := `
		    INSERT INTO sportsman
		    	(name,surname,phone,email,password,time_create,role_uuid)
		    VALUES
				($1,$2,$3,$4,$5,$6,$7)
		    RETURNING sportsman_uuid
				`

	if err := tx.QueryRow(ctx, q, sportsman.Name, sportsman.Surname, sportsman.Phone, sportsman.Email, sportsman.Password, sportsman.Time_create, sportsman.Role_uuid).Scan(&sportsman.Sportsman_uuid); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return "", newErr
		}
		return "", err
	}
	return sportsman.Sportsman_uuid, nil

}

func (r *repositoryAuthSportsman) GetSportsman(ctx context.Context, sportsman modelSportsman.SignInDTO) (modelSportsman.AuthDTO, error) {

	tx, err := r.client.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return modelSportsman.AuthDTO{}, err
	}
	defer tx.Commit(ctx)

	q := `
			SELECT sportsman_uuid, role_uuid, email
			FROM sportsman
			WHERE email=$1 AND password=$2
				`
	var sportsmanAuth modelSportsman.AuthDTO

	err = tx.QueryRow(ctx, q, sportsman.Email, sportsman.Password).Scan(&sportsmanAuth.Sportsman_uuid, &sportsmanAuth.Role, &sportsmanAuth.Email)
	if err != nil {
		tx.Rollback(ctx)
		return modelSportsman.AuthDTO{}, err
	}

	role := `
				SELECT name
				FROM role
				WHERE role_uuid=$1
			`
	err = tx.QueryRow(ctx, role, sportsmanAuth.Role).Scan(&sportsmanAuth.Role)
	if err != nil {
		tx.Rollback(ctx)
		return modelSportsman.AuthDTO{}, err
	}
	return sportsmanAuth, nil
}
