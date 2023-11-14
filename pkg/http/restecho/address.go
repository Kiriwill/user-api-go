package restecho

import (
	"database/sql"
	"net/http"

	"github.com/kiriwill/desafio-verifymy/pkg/service"
	"github.com/labstack/echo/v4"
)

func SearchAddressHandler(svc service.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		address_id := c.Param("id")
		if address_id == "" {
			return c.NoContent(http.StatusBadRequest)
		}
		address, err := svc.LookupAddressById(ctx, address_id)
		if err == sql.ErrNoRows {
			return c.NoContent(http.StatusBadRequest)
		}
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, address)
	}
}

func UpdateAddressHandler(svc service.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		address_id := c.Param("id")
		if address_id == "" {
			return c.NoContent(http.StatusBadRequest)
		}

		var err error
		ctx := c.Request().Context()

		req := new(service.Address)
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		if err := c.Validate(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		r, err := svc.UpdateAddress(ctx, req, address_id)
		if err != nil {
			return err
		}
		if r == 0 {
			return c.NoContent(http.StatusOK)
		}

		return c.JSON(http.StatusOK, "updated")
	}
}

func DeleteAddressHandler(svc service.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		address_id := c.Param("id")
		if address_id == "" {
			return c.NoContent(http.StatusBadRequest)
		}

		var err error
		ctx := c.Request().Context()

		r, err := svc.DeleteAddress(ctx, address_id)
		if err != nil {
			return err
		}
		if r == 0 {
			return c.String(http.StatusOK, "No change.")
		}

		return c.String(http.StatusOK, "Address removed")
	}
}
