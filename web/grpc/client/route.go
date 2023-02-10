package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/hello", func(c *gin.Context) {
		var request HelloRequest
		err := c.ShouldBindQuery(&request)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("request" + request.Name)
		result := sayHello(request.Name)
		resp := Response{Code: 0, Msg: "success", Data: result}
		c.JSON(200, &resp)
	})
	return r
}

type HelloRequest struct {
	Name string `form:"name" binding:"required"`
}
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
