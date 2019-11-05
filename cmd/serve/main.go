package main

import (
	log "github.com/blockchain-abstraction-middleware/project-template/pkg/logger"
	"github.com/blockchain-abstraction-middleware/project-template/pkg/routes"
	"github.com/blockchain-abstraction-middleware/project-template/pkg/server"
)

func main() {
	serverConfig := server.Config{
		BasePath:       "/api/v1",
		Name:           "project-template",
		Port:           8080,
		MetricsEnabled: false,
	}
	srv := server.New(&serverConfig)

	healthResource := routes.HealthResource{}
	swaggerResource := routes.SwaggerResource{}

	srv.RegisterResource(healthResource.NewResource("/health"))
	srv.RegisterResource(swaggerResource.NewResource("/swagger"))

	if err := srv.Run(); err != nil {
		log.WithError(err).Fatal("Serving failed")
	}
}
