package controller

import (
	"fmt"

	"gocrud/model"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err == nil {
		fmt.Printf("user obj - %+v \n", user)
	} else {
		fmt.Printf("error - %+v \n", err)
	}

	c.JSON(201, user)

}
