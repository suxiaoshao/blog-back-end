package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type LoginClaims struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// Encryption 编码
func Encryption(userName string, password string, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, LoginClaims{
		UserName: userName,
		Password: password,
		StandardClaims: jwt.StandardClaims{},
	})
	// SecretKey 用于对用户数据进行签名，不能暴露
	return token.SignedString([]byte(secretKey))
}

// Decryption 解码
func Decryption(tokenString string,secretKey string) (*LoginClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &LoginClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*LoginClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}