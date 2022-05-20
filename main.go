package main

import (
	"github.com/aeof/douyin/api"
	_ "github.com/aeof/douyin/model"
	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()

	// Unauthorized APIs
	{
		c.POST("douyin/user/register/", api.Register)
		c.POST("douyin/user/login/", api.Login)
	}

	// Authorized APIs all contain a query parameter
	authGroup := c.Group("/douyin")
	authGroup.Use(api.AuthMiddleware)
	{
		authGroup.GET("/user/", api.QueryUserInfo)
	}

	c.Run()
}
