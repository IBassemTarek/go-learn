package middleware

import (
	service "go-learn/services"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.NewJWTService().ValidateToken(tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[Issued At]: ", claims["iat"])
			log.Println("Claims[Expires At]: ", claims["exp"])
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
