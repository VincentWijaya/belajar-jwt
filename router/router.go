package router

import (
	"belajar-jwt/controllers"
	"belajar-jwt/middlewares"

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
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productID", middlewares.ProductAuthorization(), controllers.UpdateProduct)
	}

	return r
}
