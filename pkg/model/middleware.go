package middleware

import "net/http"

// Middleware type defines nested HandlerFuncs
type Middleware func(http.HandlerFunc) http.HandlerFunc
