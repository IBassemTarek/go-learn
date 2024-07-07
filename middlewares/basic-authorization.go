package middleware

import "github.com/gin-gonic/gin"

func BasicAuthorization() gin.HandlerFunc {
	// return the basic auth middleware
	// revieve the username and password
	return gin.BasicAuth(gin.Accounts{
		"userName": "password",
	})
}
