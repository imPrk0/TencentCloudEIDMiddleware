package routers

import (
	"TencentCloudEIDMiddleware/routers/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controllers.EidDecode)

	return r
}
