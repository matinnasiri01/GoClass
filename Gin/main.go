package main

import "github.com/gin-gonic/gin"

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	engine := gin.Default()
	engine.GET("/ping", ping)
	engine.Run()
}
