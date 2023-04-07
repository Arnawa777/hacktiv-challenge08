package controllers

import (
	"challenge-08/database"
	"challenge-08/helpers"
	"challenge-08/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	db := database.GetDB()
	user := models.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         user.ID,
		"first_name": user.FirstName,
		"email":      user.Email,
		"isAdmin":    user.IsAdmin,
	})
}

func LoginUser(ctx *gin.Context) {
	db := database.GetDB()
	user := models.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	password := user.Password

	err = db.Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if !helpers.PasswordValid(user.Password, password) {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid password"))
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}
