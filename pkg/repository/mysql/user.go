package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kiriwill/desafio-verifymy/pkg/service"
)

var (
	lookupUserQuery = `
		SELECT id, name, birthdate, email, password
		FROM customer
		WHERE email = ?;
	`

	lookupUserByIdQuery = `
		SELECT id, name, birthdate, email
		FROM customer
		WHERE id = ?;
	`

	deleteUserQuery = `
		DELETE FROM customer
		WHERE id=?;
	`
)

func (repo *MysqlRepository) DeleteUser(ctx context.Context, user_id string) (int64, error) {
	r, err := repo.db.Exec(deleteUserQuery, user_id)
	if err != nil {
		return 0, fmt.Errorf("could not update user %w", err)
	}

	return r.RowsAffected()
}

func (repo *MysqlRepository) LookupUser(ctx context.Context, username string) (*service.User, error) {
	var user = new(service.User)
	err := repo.db.QueryRow(lookupUserQuery, username).Scan(&user.ID, &user.Name, &user.Birthdate, &user.Email, &user.Password)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return user, nil
}

func (repo *MysqlRepository) LookupUserById(ctx context.Context, user_id string) (*service.User, error) {
	var user = new(service.User)
	err := repo.db.QueryRow(lookupUserByIdQuery, user_id).Scan(&user.ID, &user.Name, &user.Birthdate, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
