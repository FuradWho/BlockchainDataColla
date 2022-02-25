package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Test api for test
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, "test success")
}

// ConnNet Connect to main net
func ConnNet(c *gin.Context) {

}


