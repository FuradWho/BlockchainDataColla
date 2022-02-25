package routers

import (
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/web/controllers"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	// Api Group
	userGroup := r.Group("/test")
	{
		// Test api
		userGroup.GET("", controllers.Test)
	}
	return r
}
