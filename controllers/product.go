package controllers

import (
	"challenge-08/database"
	"challenge-08/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func GetAllProducts(ctx *gin.Context) {
	db := database.GetDB()
	products := []models.Product{}

	err := db.Find(&products).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func GetProduct(ctx *gin.Context) {
	db := database.GetDB()
	// userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = db.Where("id = ?", productID).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Product not found",
			})
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	product.UserID = uint(userData["id"].(float64))

	err = db.Create(&product).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	product.UserID = uint(userData["id"].(float64))

	err = db.Model(&product).Where("id=?", productID).Updates(models.Product{Title: product.Title, Description: product.Description}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func DeleteProduct(ctx *gin.Context) {
	db := database.GetDB()
	productID, _ := strconv.Atoi(ctx.Param("productID"))
	product := models.Product{}

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Delete the product from the database
	if err := db.Delete(&models.Product{}, productID).Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
