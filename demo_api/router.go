package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer(port int) {
	r := gin.New() // New Engine
	r.Use(gin.Recovery())
	r.GET("/", hello)
	r.GET("/users", getUsers)
	r.Run( fmt.Sprintf(":%v", port))
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Title     string `json:"title"`
}

type Users []User

func getUsers(c *gin.Context) {
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

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world")
}