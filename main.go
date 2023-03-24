package main

import (
	"gocrud/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := SetUpRouter()
	router.Run(":8081")

}

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	v1Book := router.Group("api/v1/books")
	{
		v1Book.POST("", controller.CreateBook)
	}

	v1User := router.Group("api/v1/users")
	{
		v1User.POST("", controller.CreateUser)
	}

	return router
}
