package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"
	"github.com/puwarinnueng/gin-marketplace/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/", controllers.Get )
	r.GET("/data", controllers.Data )
	r.Run(":8080")
}

