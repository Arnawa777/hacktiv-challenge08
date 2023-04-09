package router

import (
	"challenge-08/controllers"
	"challenge-08/middlewares"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}

	productRouter := r.Group("products")
	productRouter.Use(middlewares.Authentication())
	{
		productRouter.GET("/", controllers.GetAllProducts)
		productRouter.GET("/:productID", middlewares.ValidateUser(), controllers.GetProduct)
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productID", middlewares.ValidateUser(), controllers.UpdateProduct)
		productRouter.DELETE("/:productID", middlewares.ValidateUser(), controllers.DeleteProduct)
	}

	return r
}
