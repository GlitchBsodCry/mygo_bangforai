package utils

import (
	//stderr "errors"
	"time"

	"mygo_bangforai/pkg/config"
	"mygo_bangforai/pkg/interfacer"
	"mygo_bangforai/api/errors"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

var logger = interfacer.GetLogger()

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

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", errors.WrapError(err, errors.UtilsError, "生成JWT失败", "internal/utils/jwt.go/GenerateToken")
	}
	return token, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, errors.WrapError(err, errors.UtilsError, "解析JWT失败", "internal/utils/jwt.go/ParseToken")
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.NewError(errors.UtilsError, "invalid token", "internal/utils/jwt.go/ParseToken")
}