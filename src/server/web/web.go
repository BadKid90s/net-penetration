package main

import (
	"github.com/gin-gonic/gin"
	"net-penetration/define"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {

		c.String(http.StatusOK, "Hello World")

	})

	r.Run(define.LocalServerAddress)

}
