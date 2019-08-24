package hello

import (
	"fmt"
	"net/http"

	middleware "github.com/blockchain-abstraction-middleware/project-template/model"
	"github.com/gorilla/schema"
)

type Payload struct {
	hello string
}

// Hello ::POST prints the req body
func Hello() middleware.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseForm()

			p := new(Payload)
			decoder := schema.NewDecoder()

			fmt.Println(r.PostForm)

			err = decoder.Decode(p, r.Form)

			if err != nil {
				fmt.Println("Error decoding")
			}

			fmt.Fprintln(w, "202")

			f(w, r)
		}
	}
}
