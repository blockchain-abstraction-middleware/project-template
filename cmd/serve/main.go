package main

import (
	log "github.com/blockchain-abstraction-middleware/rest-api/pkg/logger"
	"github.com/blockchain-abstraction-middleware/rest-api/pkg/routes"
	"github.com/blockchain-abstraction-middleware/rest-api/pkg/server"
	// config "github.com/blockchain-abstraction-middleware/project-template/pkg/config"
)

func main() {
	// abstractionConfig := config.NewConfig()

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
