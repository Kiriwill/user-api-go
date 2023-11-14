package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kiriwill/desafio-verifymy/pkg/service"
)

var (
	createUserQuery = `
		INSERT INTO customer
		(name, birthdate, email, password)
		VALUES(?, ?, ?, ?);
	`

	createAddressQuery = `
		INSERT INTO address
		(user_id, street_address, street_number, city, region, postal_code, country)
		VALUES(?, ?, ?, ?, ?, ?, ?);
	`

	updateUserQuery = `
		UPDATE customer
		SET name=?, birthdate=?, email=?, password=?
		WHERE id=?;
	`
)

func (repo *MysqlRepository) CreateUser(tx *sql.Tx, user *service.User) (sql.Result, error) {
	r, err := tx.Exec(createUserQuery, user.Name, user.Birthdate, user.Email, user.Password)
	if err != nil {
		return nil, fmt.Errorf("could not create user %w", err)
	}

	return r, nil
}

func (repo *MysqlRepository) CreateAddress(tx *sql.Tx, user_id int64, address *service.Address) error {
	_, err := tx.Exec(createAddressQuery, user_id, address.Street_name, address.Street_number, address.City, address.Region, address.Postal_code, address.Country)
	if err != nil {
		return fmt.Errorf("could not address user %w", err)
	}

	return nil
}

func (repo *MysqlRepository) CreateUserAddress(ctx context.Context, user *service.User) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("db transaction canceled due %w", err)
	}
	defer tx.Rollback()

	newUser, err := repo.CreateUser(tx, user)
	if err != nil {
		return err
	}

	user_id, err := newUser.LastInsertId()
	if err != nil {
		return err
	}

	for _, address := range user.Address {
		if err := repo.CreateAddress(tx, user_id, &address); err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo *MysqlRepository) UpdateUser(tx *sql.Tx, user *service.User, user_id string) error {
	_, err := tx.Exec(updateUserQuery, user.Name, user.Birthdate, user.Email, user.Password, user_id)
	if err != nil {
		return fmt.Errorf("could not update user %w", err)
	}

	return nil
}

func (repo *MysqlRepository) UpdateUserAddress(tx *sql.Tx, user_id int, address *service.Address) error {
	_, err := tx.Exec(updateAddressFromUserQuery, user_id, address.Street_name, address.Street_number,
		address.City, address.Region, address.Postal_code, address.Country, address.ID)
	if err != nil {
		return fmt.Errorf("could not address user %w", err)
	}

	return nil
}

func (repo *MysqlRepository) UpdateUserAddresses(ctx context.Context, user *service.User, user_id string) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("db transaction canceled due %w", err)
	}
	defer tx.Rollback()

	err = repo.UpdateUser(tx, user, user_id)
	if err != nil {
		return err
	}

	for _, address := range user.Address {

		if err := repo.UpdateUserAddress(tx, user.ID, &address); err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
