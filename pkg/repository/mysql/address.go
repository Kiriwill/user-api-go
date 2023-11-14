package mysql

import (
	"context"
	"fmt"

	"github.com/kiriwill/desafio-verifymy/pkg/service"
)

var (
	lookupAddressByIdQuery = `
		SELECT id, user_id, street_address, street_number, city, region, postal_code, country
		FROM address
		WHERE id = ?;
	`

	updateAddressFromUserQuery = `
		UPDATE address
		SET user_id=?, street_address=?, street_number=?, city=?, region=?, postal_code=?, country=?
		WHERE id=?;
	`

	updateAddressQuery = `
		UPDATE address
		SET street_address=?, street_number=?, city=?, region=?, postal_code=?, country=?
		WHERE id=?;
	`
	deleteAddressQuery = `
		DELETE FROM address
		WHERE id=?;
	`
)

func (repo *MysqlRepository) DeleteAddress(ctx context.Context, address_id string) (int64, error) {
	r, err := repo.db.Exec(deleteAddressQuery, address_id)
	if err != nil {
		return 0, fmt.Errorf("could not update user %w", err)
	}

	return r.RowsAffected()
}

func (repo *MysqlRepository) LookupAddressById(ctx context.Context, address_id string) (*service.Address, error) {
	var address = new(service.Address)
	err := repo.db.QueryRow(lookupAddressByIdQuery, address_id).Scan(&address.ID, &address.User_id, &address.Street_name, &address.Street_number, &address.City, &address.Region, &address.Postal_code, &address.Country)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (repo *MysqlRepository) UpdateAddressFromUser(ctx context.Context, address *service.Address, address_id string) (int64, error) {
	r, err := repo.db.Exec(updateAddressFromUserQuery, address.User_id, address.Street_name, address.Street_number, address.City, address.Region, address.Postal_code, address.Country, address_id)
	if err != nil {
		return 0, fmt.Errorf("could not update address %w", err)
	}

	return r.RowsAffected()
}

func (repo *MysqlRepository) UpdateAddress(ctx context.Context, address *service.Address, address_id string) (int64, error) {
	r, err := repo.db.Exec(updateAddressQuery, address.Street_name, address.Street_number, address.City, address.Region, address.Postal_code, address.Country, address_id)
	if err != nil {
		return 0, fmt.Errorf("could not update address %w", err)
	}

	return r.RowsAffected()
}
