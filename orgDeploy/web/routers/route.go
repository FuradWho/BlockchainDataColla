package routers

import (
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/web/controllers"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	// Api Group
	testGroup := r.Group("/test")
	{
		// Test apiserver
		testGroup.GET("", controllers.Test)
	}

	userGroup := r.Group("/user")
	{
		userGroup.POST("login", controllers.Login)
	}

	return r
}
