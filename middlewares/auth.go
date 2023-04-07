package middlewares

import (
	"challenge-08/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			return
		}

		c.Set("userData", claims)

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is an admin

		datas := c.MustGet("userData").(jwt.MapClaims)
		// isAdmin :=

		isAdmin, ok := datas["isAdmin"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "User role not found",
			})
			return
		}
		if isAdmin != true {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized.. User is not a Administrator",
			})
			return
		}

		// Call the next middleware function
		c.Next()
	}
}

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is an admin

		datas := c.MustGet("userData").(jwt.MapClaims)
		isAdmin, ok := datas["isAdmin"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "User role not found",
			})
			return
		}
		if isAdmin == true {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "User is a Administrator.. Please Change to Admin route",
			})
			return
		}

		// Call the next middleware function
		c.Next()
	}
}
