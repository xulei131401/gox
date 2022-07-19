package jwt

import (
	jwtGo "github.com/golang-jwt/jwt/v4"
)

//GetToken 根据参数获取jwt token
func GetToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwtGo.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwtGo.New(jwtGo.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
