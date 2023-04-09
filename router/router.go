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
		productRouter.GET("/:productID", controllers.GetProduct, middlewares.ValidateUser())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productID", controllers.UpdateProduct, middlewares.ValidateUser())
		productRouter.DELETE("/:productID", controllers.DeleteProduct, middlewares.ValidateUser())
	}

	return r
}
