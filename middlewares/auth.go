package middlewares

import (
	"challenge-08/database"
	"challenge-08/helpers"
	"challenge-08/models"
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
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		// Check if the user is an admin
		isAdmin := userData["is_admin"].(bool)
		user := models.User{}

		err := db.Select("user_id").First(&user, isAdmin).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unauthorized",
				"error":   "Failed to find product",
			})
			return
		}

		if isAdmin == false {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
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
		isAdmin, ok := c.Get("isAdmin")
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "User role not found",
			})
			return
		}
		if isAdmin.(bool) == true {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		// Call the next middleware function
		c.Next()
	}
}
