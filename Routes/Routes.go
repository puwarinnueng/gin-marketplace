package Routes

import (
	"github.com/puwarinnueng/gin-marketplace/Controllers"
	"github.com/puwarinnueng/gin-marketplace/Pages"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/get", Pages.Get)
	grp1 := r.Group("/api")
	{
		grp1.GET("user", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
	}



	return r
}
