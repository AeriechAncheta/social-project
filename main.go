package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// check if the server is running
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!!!")
	})

	// group: api v1
	// {
	// 	v1 := router.Group("/api/v1")

	// }

	router.Run(":8080")
}
