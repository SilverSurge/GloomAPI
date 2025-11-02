package bloomapi

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {

	addChansToPool()

	router.GET("/ping", pingHandler)

	router.GET("/stats", statsHandler)

	router.POST("/filters", createFilterHandler)

	router.GET("/filters", listFiltersHandler)

	router.POST("/filters/:id/add", addElementsHandler)

	router.POST("/filters/:id/check", checkElementsHandler)

}
