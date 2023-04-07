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

	isAdminRouter := r.Group("isAdmin")
	isAdminRouter.Use(middlewares.Authentication())
	isAdminRouter.Use(middlewares.AdminMiddleware())
	{
		isAdminRouter.GET("/products", controllers.GetAllProducts)
		isAdminRouter.GET("/products/:productID", controllers.GetProduct)
		isAdminRouter.POST("/products", controllers.CreateProduct)
		isAdminRouter.PUT("/products/:productID", controllers.UpdateProduct)
		isAdminRouter.DELETE("/products/:productID", controllers.DeleteProduct)
	}

	isUserRouter := r.Group("isUser")
	isUserRouter.Use(middlewares.Authentication())
	isUserRouter.Use(middlewares.UserMiddleware())
	{
		isUserRouter.GET("/products/:productID", controllers.GetProduct)
		isUserRouter.POST("/products", controllers.CreateProduct)
	}

	return r
}
