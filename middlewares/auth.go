package middlewares

import (
	"challenge-08/database"
	"challenge-08/helpers"
	"challenge-08/models"
	"net/http"
	"strconv"

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

func ValidateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productID, err := strconv.Atoi(c.Param("productID"))
		product := models.Product{}

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Unauthorized",
				"error":   "Invalid product ID data type",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		// if !ok {
		// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		// 		"message": "User id not found",
		// 	})
		// 	return
		// }

		admin, ok := userData["isAdmin"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "User role not found",
			})
			return
		}

		//If ADMIN JUST GO
		//Code stop here when true
		if admin == true {
			c.Next()
			return
		}

		err = db.Select("user_id").First(&product, uint(productID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unauthorized",
				"error":   "Failed to find product",
			})
			return
		}

		if product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
				"error":   "You are not allowed to access this product",
			})
			return
		}

		c.Next()
	}
}
