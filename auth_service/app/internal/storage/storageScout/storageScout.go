package storageScout

import (
	"context"
	"fmt"
	"github.com/GermanBogatov/auth_service/internal/model"
	"github.com/GermanBogatov/auth_service/internal/model/modelScout"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/GermanBogatov/auth_service/pkg/postgresql"
	"github.com/jackc/pgconn"
)

const (
	roleScout = "scout"
)

type repositoryAuthScout struct {
	client postgresql.ClientPostgres
	logger logging.Logger
}

func NewAuthorization(client postgresql.ClientPostgres, logger logging.Logger) *repositoryAuthScout {
	return &repositoryAuthScout{
		client: client,
		logger: logger,
	}
}

func (r *repositoryAuthScout) CreateScout(ctx context.Context, scout modelScout.ScoutDTO) (string, error) {

	tx, err := r.client.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}
	defer tx.Commit(ctx)
	var roleUUID string
	role := `
				SELECT role_uuid
				FROM Roles
				WHERE name=$1
			`
	if err := tx.QueryRow(ctx, role, roleScout).Scan(&roleUUID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return "", newErr
		}
		return "", err
	}

	q := `
			    INSERT INTO Scout
			    	(name,surname,phone,email,password,time_create,role_uuid)
			    VALUES
					($1,$2,$3,$4,$5,$6,$7)
			    RETURNING scout_uuid
					`

	if err := tx.QueryRow(ctx, q, scout.Name, scout.Surname, scout.Phone, scout.Email, scout.Password, scout.Time_create, roleUUID).Scan(&scout.Scout_uuid); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return "", newErr
		}
		return "", err
	}
	return scout.Scout_uuid, nil

}

func (r *repositoryAuthScout) SignInScout(ctx context.Context, scout model.SignInDTO) (model.AuthDTO, error) {

	q := `
	SELECT
		sc.scout_uuid, sc.email, ro.name
	FROM
		Scout sc
	INNER JOIN
		Roles ro on ro.role_uuid = sc.role_uuid
	WHERE 
		sc.email = $1
	AND
		sc.password = $2
		`
	var scoutAuth model.AuthDTO

	err := r.client.QueryRow(ctx, q, scout.Email, scout.Password).Scan(&scoutAuth.Uuid, &scoutAuth.Email, &scoutAuth.Role)
	if err != nil {
		return model.AuthDTO{}, err
	}

	return scoutAuth, nil
}
