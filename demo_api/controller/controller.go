package controller

import (
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	us := service.UserService{}
	c.JSON(http.StatusOK, us.GetUsers())
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world")
}