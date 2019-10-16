package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/blockchain-abstraction-middleware/project-template/pkg/server"
	"github.com/go-chi/chi"
)

// HealthResource implements server.Resource interface
type HealthResource struct {
	path string
}

// NewResource constructor func
func (r *HealthResource) NewResource(path string) server.Resource {
	return &HealthResource{
		path: path,
	}
}

// Path returns the Resource base path
func (r *HealthResource) Path() string {
	return r.path
}

// Routes bootstraps health routes
func (r *HealthResource) Routes() http.Handler {
	router := chi.NewRouter()

	router.Get("/status", r.healthCheck())

	return router
}

func (r *HealthResource) healthCheck() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		payload := struct {
			Message    string `json:"message"`
			StatusCode int    `json:"statusCode"`
		}{
			fmt.Sprintf("Healthy response from service at - %s", time.Now()),
			200,
		}

		json, _ := json.Marshal(payload)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(json)
	}
}
