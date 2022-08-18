// back-end/cmd/api/main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ServerPort = "8080"
const ServerAddr = "localhost"

func main() {
	router := gin.Default()

	Addr := fmt.Sprintf("%s:%s", ServerAddr, ServerPort)
	fmt.Printf("[Backend]: Service starting on port %s\n", ServerPort)

	// Handle the incoming request

	// API ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// API status
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// Start server
	router.Run(Addr)

}
