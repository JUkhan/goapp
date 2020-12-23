package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	JWTService interface {
		GenerateToken(name string, admin bool) string
		ValidateToken(tokenString string) (*jwt.Token, error)
	}
	jwtCustomClaims struct {
		Name  string `json:"name"`
		Admin bool   `json:"admin"`
		jwt.StandardClaims
	}
	jwtService struct {
		secretKey string
		issuer    string
	}
)

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "jukhan.com",
	}
}
func getSecretKey() string {
	sk := os.Getenv("JWT_SECRET")
	if sk == "" {
		sk = "secret"
	}
	return sk
}

func (jwtSrv *jwtService) GenerateToken(name string, admin bool) string {
	claims := &jwtCustomClaims{
		name, admin, jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	//create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//create encoded token using the secret signing key
	t, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		//signing method validation
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSrv.secretKey), nil
	})
}
