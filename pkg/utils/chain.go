package utils

import (
	"net/http"

	middleware "github.com/blockchain-abstraction-middleware/contract-data/pkg/model"
)

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...middleware.Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
