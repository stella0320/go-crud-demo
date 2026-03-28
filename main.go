package main

import (
	"go-crud-demo/handler"
	"go-crud-demo/middleware"
	"go-crud-demo/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // server 路由器
	r.Use(middleware.ErrorHandler())
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)
	// route
	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.AddUser)
	r.GET("/users/:id", userHandler.QueryUserById)
	r.Run()
}
