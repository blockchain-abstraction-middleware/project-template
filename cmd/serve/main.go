package main

import (
	log "github.com/blockchain-abstraction-middleware/project-template/pkg/logger"
	"github.com/blockchain-abstraction-middleware/project-template/pkg/routes"
	"github.com/blockchain-abstraction-middleware/project-template/pkg/server"
)

func main() {
	srv := server.New(server.NewConfig())
	srv.RegisterResource(routes.NewResource("/health"))

	if err := srv.Run(); err != nil {
		log.WithError(err).Fatal("Serving failed")
	}
}
