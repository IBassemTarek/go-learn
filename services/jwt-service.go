package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(name string, admin bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

// jwt custom claims are custom claims extending default ones
// properties not included in jwt.StandardClaims ( JWT token properties )
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getScretKey(),
		issuer:    "myapp",
	}
}

func getScretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtService) GenerateToken(name string, admin bool) string {
	//  set custom and standard claims
	claims := &jwtCustomClaims{
		name,
		admin,
		jwt.StandardClaims{
			Issuer:    service.issuer,
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate encoded token using the secret key
	t, _ := token.SignedString([]byte(service.secretKey))
	return t
}

func (service *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	// parse token using the secret key
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// check token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}
