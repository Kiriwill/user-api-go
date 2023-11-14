package service

import (
	"context"
	"database/sql"
)

type (
	User struct {
		ID        int       `json:"id" validate:"omitempty"`
		Name      string    `json:"name" validate:"required"`
		Birthdate string    `json:"birthdate" validate:"required,datetime=2006-01-02"`
		Email     string    `json:"email" validate:"required,email"`
		Password  string    `json:"password" validate:"required,omitempty"`
		Address   []Address `json:"address" validate:"required,dive,required"`
	}

	Address struct {
		ID            int    `json:"id"`
		Street_name   string `json:"street_name" validate:"required"`
		Street_number string `json:"street_number" validate:"required"`
		City          string `json:"city" validate:"required"`
		User_id       int    `json:"user_id" validate:"omitempty"`
		Region        string `json:"region" validate:"required"`
		Postal_code   string `json:"postal_code"`
		Country       string `json:"country" validate:"required"`
	}

	Repository interface {
		// User
		CreateUserAddress(ctx context.Context, user *User) error
		UpdateUser(tx *sql.Tx, user *User, user_id string) error
		CreateUser(tx *sql.Tx, user *User) (sql.Result, error)
		LookupUser(ctx context.Context, query string) (*User, error)
		LookupUserById(ctx context.Context, query string) (*User, error)
		DeleteUser(ctx context.Context, user_id string) (int64, error)

		// Address
		UpdateAddress(ctx context.Context, address *Address, user_id string) (int64, error)
		UpdateUserAddress(tx *sql.Tx, user_id int, address *Address) error
		UpdateUserAddresses(ctx context.Context, user *User, user_id string) error
		CreateAddress(tx *sql.Tx, user_id int64, address *Address) error
		LookupAddressById(ctx context.Context, query string) (*Address, error)
		DeleteAddress(ctx context.Context, user_id string) (int64, error)
	}
)
