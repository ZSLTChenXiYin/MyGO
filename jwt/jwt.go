package jwt

import (
	"time"

	"github.com/ZSLTChenXiYin/MyGO/configure"
	"github.com/golang-jwt/jwt"
)

const (
	JWT_KEY_EXPIRED_AT = "expired_at"
)

type JWTGenerator struct {
	secret string
}

func NewJWTGenerator(conf configure.Configuration) *JWTGenerator {
	return &JWTGenerator{
		secret: conf.Server().JwtSecret(),
	}
}

func (j *JWTGenerator) Create(claims jwt.MapClaims) (string, error) {
	claims[JWT_KEY_EXPIRED_AT] = time.Now().Add(time.Hour * 24 * 7).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &claims).SignedString([]byte(j.secret))
}

func (j *JWTGenerator) Parse(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (j *JWTGenerator) Refresh(token string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return "", err
	}
	claims[JWT_KEY_EXPIRED_AT] = time.Now().Add(time.Hour * 24 * 7).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &claims).SignedString([]byte(j.secret))
}
