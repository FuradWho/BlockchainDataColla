package controllers

import (
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/service/fabric_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Test apiserver for test
func Test(c *gin.Context) {

	fabric_service.Msg()

	c.JSON(http.StatusOK, "test success")

}

func Login(c *gin.Context) {

}

// ConnNet Connect to main net
func ConnNet(c *gin.Context) {

}
