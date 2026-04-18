package jwt

import (
	"maps"
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

func (j *JWTGenerator) Create(map_claims map[string]any) (string, error) {
	claims := jwt.MapClaims{}
	maps.Copy(claims, map_claims)
	claims[JWT_KEY_EXPIRED_AT] = time.Now().Add(time.Hour * 24 * 7).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.secret)
}

func (j *JWTGenerator) Parse(token string) (map[string]any, error) {
	map_claims := map[string]any{}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}
	maps.Copy(map_claims, claims)
	return map_claims, nil
}

func (j *JWTGenerator) Refresh(token string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return j.secret, nil
	})
	if err != nil {
		return "", err
	}
	claims[JWT_KEY_EXPIRED_AT] = time.Now().Add(time.Hour * 24 * 7).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.secret)
}
