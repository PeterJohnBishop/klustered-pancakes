// Package main starts a Gin server
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}

		greeting := fmt.Sprintf("Your server today is, %s", hostname)

		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Klustered-Pancake!",
			"server":  greeting,
		})
	})
	r.Run(":8080")
}
