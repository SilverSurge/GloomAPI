package bloomapi

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {

	router.GET("/ping", pingHandler)

	router.POST("/filters", createFilterHandler)

	router.GET("/filters", listFiltersHandler)

}
