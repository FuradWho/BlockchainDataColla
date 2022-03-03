package controllers

import (
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/service/fabric_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Test api for test
func Test(c *gin.Context) {

	fabric_service.Msg()

	c.JSON(http.StatusOK, "test success")

}

// ConnNet Connect to main net
func ConnNet(c *gin.Context) {

}


