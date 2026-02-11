package utils

import (
	"errors"
	"time"

	"mygo_bangforai/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

// InitJWT 初始化JWT配置
func InitJWT() error{
	jwtConfig := config.GetJWTConfig()
	jwtSecret = []byte(jwtConfig.Secret)
	return nil
}

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims	//标准字段，含过期时间、签发时间等标准JWT声明
}

func GenerateToken(userID uint, username string) (string, error) {
	jwtConfig := config.GetJWTConfig()
	
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtConfig.ExpiresIn) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),//签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),//生效时间
			Issuer:    jwtConfig.Issuer,
			Subject:   jwtConfig.Subject,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}