package bloomapi

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {

	router.GET("/ping", pingHandler)

}
