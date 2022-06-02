package jwt

import (
	"github.com/ArseniySavin/auth-small-server/typos"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Make -
func (j *DefaultJwtServices) Make(claims map[string]interface{}) (string, error) {
	if claims == nil {
		claims = make(map[string]interface{})
	}

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["token_type"] = "Bearer"

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))

	return jwtToken.SignedString([]byte(j.secret))
}

// Parse -
func (j *DefaultJwtServices) Parse(rawJwt string) (map[string]interface{}, error) {
	if rawJwt == "" {
		return nil, typos.ErrEmptyJwt
	}

	token, err := jwt.Parse(rawJwt, keyFunc(j.secret))
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	if exp, ok := claims["exp"]; ok {
		claims["exp"] = int64(exp.(float64))
	}

	return claims, nil
}

func keyFunc(secret string) jwt.Keyfunc {
	return func(*jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}
}
