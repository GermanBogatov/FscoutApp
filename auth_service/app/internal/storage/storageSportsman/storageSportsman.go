package storageSportsman

import (
	"context"
	"fmt"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/GermanBogatov/auth_service/pkg/postgresql"
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

func (r *repositoryAuthSportsman) CreateSportsman(ctx context.Context) (int, error) {
	fmt.Println("storage sportsman!")
	/*	q := `
		    INSERT INTO users
		    	(name,username,password_hash)
		    VALUES
				($1,$2,$3)
		    RETURNING id
				`

			if err := r.client.QueryRow(ctx, q, user.Name, user.Username, user.Password).Scan(&user.Id); err != nil {
				if pgErr, ok := err.(*pgconn.PgError); ok {
					newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
					r.logger.Error(newErr)
					return 0, newErr
				}
				return 0, err
			}
			return user.Id, nil*/
	return 0, nil
}

func (r *repositoryAuthSportsman) GetSportsman(ctx context.Context, username, password string) error {
	/*	var user model.UserDTO

		q := `
		SELECT id, username
		FROM users
		WHERE username=$1 AND password_hash=$2
			`

		err := r.client.QueryRow(ctx, q, username, password).Scan(&user.Id, &user.Username)
		if err != nil {
			return model.UserDTO{}, err
		}

		return user, nil*/

	return nil
}