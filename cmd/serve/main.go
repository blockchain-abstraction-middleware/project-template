package main

import (
	log "github.com/blockchain-abstraction-middleware/project-template/pkg/logger"
	"github.com/blockchain-abstraction-middleware/project-template/pkg/routes"
	"github.com/blockchain-abstraction-middleware/project-template/pkg/server"
)

func main() {
	srv := server.New(server.NewConfig())

	healthResource := routes.HealthResource{}
	swaggerResource := routes.SwaggerResource{}

	srv.RegisterResource(healthResource.NewResource("/health"))
	srv.RegisterResource(swaggerResource.NewResource("/swagger"))

	if err := srv.Run(); err != nil {
		log.WithError(err).Fatal("Serving failed")
	}
}
