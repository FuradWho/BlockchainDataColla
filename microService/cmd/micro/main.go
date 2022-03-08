package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// 1.添加consul地址
	cr := consul.NewRegistry(
		registry.Addrs("192.168.175.129:8500"))

	// 2.使用gin作为router
	router := gin.Default()
	router.GET("/user", func(c *gin.Context) {
		c.String(http.StatusOK, "user apiserver")
	})

	// 3.初始化go micro_service
	server := web.NewService(
		web.Name("productService"), // 当前微服务服务名
		web.Registry(cr),           // 注册到consul
		web.Address(":8081"),       // 端口
		web.Metadata(map[string]string{"protocol": "http"}), // 元信息
		web.Handler(router)) // 路由

	_ = server.Run()
}
