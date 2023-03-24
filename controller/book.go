package controller

import (
	"fmt"

	"gocrud/model"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err == nil {
		fmt.Printf("user obj - %+v \n", book)
	} else {
		fmt.Printf("error - %+v \n", err)
	}

	c.JSON(201, gin.H{
		"message": "성공",
	})

}
