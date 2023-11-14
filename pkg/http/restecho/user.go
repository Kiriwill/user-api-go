package restecho

import (
	"database/sql"
	"net/http"

	"github.com/kiriwill/desafio-verifymy/pkg/service"
	"github.com/labstack/echo/v4"
)

func SearchUserHandler(svc service.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		user_id := c.Param("id")
		if user_id == "" {
			return c.NoContent(http.StatusBadRequest)
		}
		user, err := svc.LookupUserById(ctx, user_id)
		if err == sql.ErrNoRows {
			return c.NoContent(http.StatusBadRequest)
		}
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, user)
	}
}

func UpdateUserHandler(svc service.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := c.Param("id")
		if user_id == "" {
			return c.NoContent(http.StatusBadRequest)
		}

		var err error
		ctx := c.Request().Context()

		req := new(service.User)
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		if err := c.Validate(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		req.Password, err = EncriptPassword(req.Password)
		if err != nil {
			return err
		}

		err = svc.UpdateUserAddresses(ctx, req, user_id)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "updated")
	}
}

func DeleteUserHandler(svc service.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := c.Param("id")
		if user_id == "" {
			return c.NoContent(http.StatusBadRequest)
		}

		var err error
		ctx := c.Request().Context()

		r, err := svc.DeleteUser(ctx, user_id)
		if err != nil {
			return err
		}
		if r == 0 {
			return c.NoContent(http.StatusBadRequest)
		}

		return c.String(http.StatusOK, "User removed")
	}
}
