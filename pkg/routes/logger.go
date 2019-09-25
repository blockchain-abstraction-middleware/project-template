package routes

import (
	"log"
	"net/http"
	"time"

	middleware "github.com/blockchain-abstraction-middleware/project-template/pkg/model"
)

// Logging logs all requests
func Logging() middleware.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			f(w, r)
		}
	}
}
