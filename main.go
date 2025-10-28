package main

import (
	bloomapi "GloomAPI/bloom_api"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Gloom API")

	// get gin router
	router := gin.Default()

	// setup router
	bloomapi.SetupRouter(router)

	// run router
	if err := router.Run(); err != nil {
		log.Fatalf("[main] failed to start server: %v", err)
	}
}
