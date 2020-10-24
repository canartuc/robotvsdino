package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	runningPort := os.Getenv("PORT")

	err := startGinServer().Run(fmt.Sprintf(":%s", string(runningPort)))
	if err != nil {
		log.Fatalf("Cannot continue, Gin server error: %v", err)
	}

}

// startGinServer setups gin server with mode and necessary parameters
// we are not writing Gin server setup directly to the main function
// to create easy integration tests
func startGinServer() *gin.Engine {
	currentEnv := os.Getenv("APIENV")
	if currentEnv == "DEV" {
		gin.SetMode(gin.DebugMode)
	} else if currentEnv == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Group the API with version
	api := r.Group("/api")
	v := api.Group("/v1")

	// Route: Index
	v.GET("/", RouteIndex)

	return r
}
