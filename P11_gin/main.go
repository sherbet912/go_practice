package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	// 基礎 routes
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome!")
	})

	router.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})

	// group routes
	// create a new group
	v1 := router.Group("v1")
	v2 := router.Group("v2")
	// Define a route for the root URL 
	v1.GET("/hello", hello_v1)		// 0.0.0.0:8080/v1/hello
	v1.GET("/hello_v2", hello_v2)	// 0.0.0.0:8080/v1/hello_v2
	v2.GET("/hello", hello_v2)		// 0.0.0.0:8080/v2/hello

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func hello_v1(c *gin.Context) {
	c.String(http.StatusOK, "Hello v1.")
}

func hello_v2(c *gin.Context) {
	c.String(http.StatusOK, "Hello v2.")
}