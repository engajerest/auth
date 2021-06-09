package app

import "github.com/engajerest/auth/controller"

// "github.com/gin-gonic/gin"

func Mapurls() {

//dev
	router.GET("/dev", controller.PlaygroundHandlers())
	router.POST("/dev/auth", controller.GraphHandler())
//live
	router.GET("/v1", controller.PlaygroundHandlers())
	router.POST("/v1/auth", controller.GraphHandler())
}
