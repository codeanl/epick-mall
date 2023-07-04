package common

import "github.com/golang-jwt/jwt/v4"

func GetJwtToken(secretKey string, iat, seconds, id int64, Phone string, nickname string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["id"] = id
	claims["phone"] = Phone
	claims["nickname"] = nickname
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
func GetInvalidJwtToken(secretKey string, iat, id int64, phone string, nickname string) (string, error) {
	exp := iat - 1 // 将过期时间设置为当前时间的前一秒
	claims := jwt.MapClaims{
		"exp":      exp,
		"iat":      iat,
		"id":       id,
		"phone":    phone,
		"nickname": nickname,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
