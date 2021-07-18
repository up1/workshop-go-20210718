package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Title     string `json:"title"`
}

type Users []User

func GetUsers(c *gin.Context) {
	u := Users{
		User{
			Firstname: "f1",
			Lastname:  "l1",
			Title:     "Mr.",
		},
		User{
			Firstname: "f2",
			Lastname:  "l2",
			Title:     "Miss.",
		},
	}
	c.JSON(http.StatusOK, u)
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world")
}