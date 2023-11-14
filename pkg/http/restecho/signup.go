package restecho

import (
	"net/http"

	"github.com/kiriwill/desafio-verifymy/pkg/service"
	"github.com/labstack/echo/v4"
)

func SignUpHandler(svc service.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
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

		err = svc.CreateUserAddress(ctx, req)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		return c.NoContent(http.StatusOK)
	}
}
