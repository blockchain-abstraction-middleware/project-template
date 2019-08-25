package routes

import (
	"fmt"
	"net/http"
)

// Status returns a status code
func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "202")
}
