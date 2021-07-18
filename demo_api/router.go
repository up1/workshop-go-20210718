package api

import (
	"api/controller"
	"fmt"

	"github.com/labstack/echo/v4"
)

func NewServer(port int) {
	r := echo.New() // New Engine
	r.GET("/", controller.Hello)
	r.GET("/users", controller.GetUsers)
	r.Start( fmt.Sprintf(":%v", port))
}

