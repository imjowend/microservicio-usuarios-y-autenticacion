package api

import (
	"github.com/gin-gonic/gin"
	hdl "github.com/imjowend/microservicio-usuarios-y-autenticacion/cmd/api/handler"
	ucs "github.com/imjowend/microservicio-usuarios-y-autenticacion/internal/core"
	"log"
)

func Build(dep *Dependencies) *gin.Engine {
	usecase := ucs.NewUseCase(dep.Repository)
	handler := hdl.NewRestHandler(usecase)
	r := gin.Default()
	v1 := r.Group("/api/v1/users")
	{
		v1.POST("/create", handler.FakeCreateUser)
	}
	log.Println("Running server on port:" + dep.RouterPort + "...")
	if err := r.Run(dep.RouterPort); err != nil {
		log.Fatalf("Failed to run server: %s\n", err)
	}
	return r
}
