package u

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sanyewudezhuzi/tiktok/conf"
)

type Claims struct {
	ID      uint   `json:"id"`
	Account string `json:"account"`
	jwt.StandardClaims
}

// GenerateToken 生成 token 令牌
func GenerateToken(uid uint, account string) (string, error) {
	issTime := time.Now()
	expTime := issTime.Add(time.Hour * 24)
	claims := Claims{
		ID:      uid,
		Account: account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
			IssuedAt:  issTime.Unix(),
			Issuer:    "tiktok",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString([]byte(conf.SecretKey))
	return tokenstr, err
}

// ParseToken 解析 token 令牌
func ParseToken(tokenstr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenstr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.SecretKey), nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
