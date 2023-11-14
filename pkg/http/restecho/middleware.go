package restecho

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type auth struct {
	secretKey []byte
	duration  time.Duration
}

type userClaims struct {
	jwt.RegisteredClaims
}

func NewAuth(secretKey string, duration time.Duration) *auth {
	return &auth{
		secretKey: []byte(secretKey),
		duration:  duration,
	}
}

func (a *auth) NewUserClaim() *userClaims {
	return &userClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * a.duration)),
		},
	}
}

func (a *auth) GenerateToken(claims *userClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(a.secretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *auth) CustomJwtMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(userClaims)
		},
		SigningKey: []byte(a.secretKey),
	})
}

// EncriptPass ensures that two users don't have the same hash
func EncriptPassword(passwd string) (string, error) {
	paswdBytes := []byte(passwd)

	hashedPass, err := bcrypt.GenerateFromPassword(paswdBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPass), err
}

func ComparePassword(hashedPasswd string, passwd string) (bool, error) {
	hashedPasswdBytes := []byte(hashedPasswd)
	passwdBytes := []byte(passwd)

	if err := bcrypt.CompareHashAndPassword(hashedPasswdBytes, passwdBytes); err != nil {
		return false, nil
	}
	return true, nil
}
