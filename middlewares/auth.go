package middlewares

import (
	"challenge-08/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helpers.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
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
		// Get the user data from the context
		userData := c.MustGet("userData").(jwt.MapClaims)

		isAdmin := userData["isAdmin"].(bool)

		// Check the user's role
		if isAdmin == true {
			c.Next()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "Only admin can access this route",
			})
			return
		}
	}
}

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user data from the context
		userData := c.MustGet("userData").(jwt.MapClaims)

		isAdmin := userData["isAdmin"].(bool)

		// Check the user's role
		if isAdmin == false {
			c.Next()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "Only users can access this route",
			})
			return
		}
	}
}
