package model

type Book struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:price binding:"required`
}
