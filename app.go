package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func get_logan(c *gin.Context) {
	c.HTML(http.StatusOK, "logan.html", gin.H{})
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", get_logan)
	gin.SetMode(gin.ReleaseMode)
	router.Run("127.0.0.1:5000")
}
