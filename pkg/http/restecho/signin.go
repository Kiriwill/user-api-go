package restecho

import (
	"database/sql"
	"net/http"

	"github.com/kiriwill/desafio-verifymy/pkg/service"
	"github.com/labstack/echo/v4"
)

type UserCredentials struct {
	Email string `json:"email" validate:"required"`
	Pass  string `json:"password" validate:"required"`
}

func SignInHandler(svc service.Repository, auth *auth) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := new(UserCredentials)
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		if err := c.Validate(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		user, err := svc.LookupUser(ctx, req.Email)
		if err == sql.ErrNoRows {
			return c.NoContent(http.StatusBadRequest)
		}
		if err != nil {
			return err
		}

		userExist, err := ComparePassword(user.Password, req.Pass)
		if err != nil {
			return err
		}

		if userExist {
			claims := auth.NewUserClaim()
			token, err := auth.GenerateToken(claims)
			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, echo.Map{
				"token": token,
			})
		}

		return c.NoContent(http.StatusBadRequest)
	}
}
