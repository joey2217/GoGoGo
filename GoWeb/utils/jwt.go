package utils

import (
	"errors"
	"math/rand"
	"time"
	"web/dao/model"

	"github.com/golang-jwt/jwt/v5"
)

const EXPIRE_TIME = time.Hour * 24 * 7

// 随机字符串
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 签名密钥
var signingKey = []byte("mmFkzpYIba")

type AppClaims struct {
	UserID uint
	Role   model.Role
	jwt.RegisteredClaims
}

func randStr(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(rand_bytes)
}

func GenerateToken(id uint) (string, error) {
	claim := AppClaims{
		UserID: id,
		Role: 0,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go", // 签发者
			Subject:   "token",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(EXPIRE_TIME)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
			ID:        randStr(10),                                     // wt ID, 类似于盐值
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(signingKey)
	return token, err
}

func ParseToken(token_string string) (*AppClaims, error) {
	token, err := jwt.ParseWithClaims(token_string, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*AppClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
