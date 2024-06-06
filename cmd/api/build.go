package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Build(dep *Dependencies) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1/")
	{
		v1.POST("/create-user", user.CreateUser)
	}
	log.Println("Running server on port:" + dep.RouterPort + "...")
	if err := r.Run(dep.RouterPort); err != nil {
		log.Fatalf("Failed to run server: %s\n", err)
	}
	return r
}
