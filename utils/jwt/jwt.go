package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ohdat/app/response"
	"github.com/spf13/viper"
)

type JwtInfo struct {
	Uid        int    `json:"uid"`
	Platform   string `json:"platform"`
	UnionID    string `json:"union_id"`
	OpenID     string `json:"open_id"`
	SessionKey string `json:"session_key"`
}

type jwtCustomClaims struct {
	jwt.StandardClaims
	// 追加自己需要的信息
	JwtInfo
}

// ParseToken Rsa256 token解码
func ParseToken(tokenString string) (*JwtInfo, error) {

	var publicKey = viper.GetString("jwt.public")
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		log.Println("ParseToken ParseRSAPublicKeyFromPEM err:", err)
		return nil, err
	}
	// https://gowalker.org/github.com/dgrijalva/jwt-go#ParseWithClaims
	// 输入用户自定义的Claims结构体对象,token,以及自定义函数来解析token字符串为jwt的Token结构体指针
	// Keyfunc是匿名函数类型: type Keyfunc func(*Token) (interface{}, error)
	// func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {}
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {

		_, err := token.Method.(*jwt.SigningMethodRSA)
		if !err {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else {
			return verifyKey, nil
		}
		//return []byte(secret), nil
	})

	if err != nil {
		// https://gowalker.org/github.com/dgrijalva/jwt-go#ValidationError
		// jwt.ValidationError 是一个无效token的错误结构
		if ve, ok := err.(*jwt.ValidationError); ok {
			// ValidationErrorMalformed是一个uint常量，表示token不可用
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, response.ErrTokenVerificationFail
				// ValidationErrorExpired表示Token过期
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				//token 过期
				return nil, response.ErrTokenExpire
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, response.ErrTokenVerificationFail
			} else {
				return nil, response.ErrTokenVerificationFail
			}
		}
	}
	// 将token中的claims信息解析出来并断言成用户自定义的有效载荷结构
	if claims, ok := token.Claims.(*jwtCustomClaims); ok && token.Valid {
		return &claims.JwtInfo, nil
	}
	return nil, fmt.Errorf("token failed ")

}

// CreateToken  使用 Rsa256 加密
func CreateToken(info JwtInfo) (tokenString string, err error) {
	var privateKey = viper.GetString("jwt.private")
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return
	}

	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "login",
		},
		info,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err = token.SignedString(signKey)
	return
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
