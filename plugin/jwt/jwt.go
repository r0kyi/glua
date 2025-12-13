package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/r0kyi/glua/core"
)

type Jwt struct {
}

func (j *Jwt) sign(key string, alg string, jwt_ map[string]any) (string, error) {
	claims := jwt.MapClaims(jwt_)
	token := jwt.NewWithClaims(getSigningMethod(alg), claims)

	raw, err := token.SignedString(core.S2B(key))
	if err != nil {
		return "", err
	}

	return raw, nil
}

func (j *Jwt) verify(key string, raw string) (map[string]any, error) {
	token, err := jwt.Parse(raw, func(_ *jwt.Token) (any, error) {
		return core.S2B(key), nil
	})
	if err != nil {
		return nil, err
	}

	jwt, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return jwt, nil
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
