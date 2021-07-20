// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	// "net/http"
// 	"github.com/puwarinnueng/gin-marketplace/controllers"
// )

// func main() {
// 	r := gin.Default()
// 	r.GET("/", controllers.Get)
// 	r.GET("/data", controllers.Data)
// 	r.Run(":8080")
// }



package main

import (
	"fmt"
	"github.com/puwarinnueng/gin-marketplace/Models"
	"github.com/puwarinnueng/gin-marketplace/Routes"
	"github.com/puwarinnueng/gin-marketplace/Config"
	"github.com/jinzhu/gorm"
	
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Employee{})
	r := Routes.SetupRouter()
	//running
	r.Run(":8080")
}
