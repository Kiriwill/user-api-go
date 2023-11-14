package restecho

import (
	"net/http"
	"time"

	"github.com/kiriwill/desafio-verifymy/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(svc service.Repository, authTokenSecret string, jwtDuration time.Duration) *echo.Echo {
	e := echo.New()
	e.Validator = NewValidator()
	e.Use(middleware.Logger())

	auth := NewAuth(authTokenSecret, jwtDuration)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	v1 := e.Group("v1")
	v1.POST("/signup", SignUpHandler(svc))
	v1.POST("/signin", SignInHandler(svc, auth))

	user := v1.Group("/user", auth.CustomJwtMiddleware())
	user.GET("/:id", SearchUserHandler(svc))
	user.PUT("/:id", UpdateUserHandler(svc))
	user.DELETE("/:id", DeleteUserHandler(svc))

	address := v1.Group("/address", auth.CustomJwtMiddleware())
	address.GET("/:id", SearchAddressHandler(svc))
	address.PUT("/:id", UpdateAddressHandler(svc))
	address.DELETE("/:id", DeleteAddressHandler(svc))

	return e
}
