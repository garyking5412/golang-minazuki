package router

import "github.com/gin-gonic/gin"

func CategoryRouter(router *gin.RouterGroup) {
	router.GET("category/all")
}
