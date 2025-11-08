package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/r0kyi/glua/core"
)

type JWT struct {
	key string
	alg string
	raw string
	jwt map[string]interface{}
}

func (j *JWT) sign() error {
	claims := jwt.MapClaims(j.jwt)
	token := jwt.NewWithClaims(getSigningMethod(j.alg), claims)

	raw, err := token.SignedString(core.S2B(j.key))
	if err != nil {
		return err
	}
	j.raw = raw

	return nil
}

func (j *JWT) verify() error {
	token, err := jwt.Parse(j.raw, func(_ *jwt.Token) (interface{}, error) {
		return core.S2B(j.key), nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token")
	}
	j.jwt = claims

	return nil
}

func getSigningMethod(alg string) jwt.SigningMethod {
	switch alg {
	case "HS256":
		return jwt.SigningMethodHS256
	case "HS384":
		return jwt.SigningMethodHS384
	case "HS512":
		return jwt.SigningMethodHS512
	case "RS256":
		return jwt.SigningMethodRS256
	case "RS384":
		return jwt.SigningMethodRS384
	case "RS512":
		return jwt.SigningMethodRS512
	case "PS256":
		return jwt.SigningMethodPS256
	case "PS384":
		return jwt.SigningMethodPS384
	case "PS512":
		return jwt.SigningMethodPS512
	case "ES256":
		return jwt.SigningMethodES256
	case "ES384":
		return jwt.SigningMethodES384
	case "ES512":
		return jwt.SigningMethodES512
	case "EdDSA":
		return jwt.SigningMethodEdDSA
	default:
		return jwt.SigningMethodHS256
	}
}
