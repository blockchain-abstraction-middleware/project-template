package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	"github.com/blockchain-abstraction-middleware/project-template/hello"
	"github.com/blockchain-abstraction-middleware/project-template/logger"
	middleware "github.com/blockchain-abstraction-middleware/project-template/model"
)

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...middleware.Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Status returns a status code
func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "202")
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	fmt.Println("Environment:", viper.Get("environment"))
	http.HandleFunc("/", Chain(Status, hello.Hello(), logger.Logging()))
	http.ListenAndServe(":8080", nil)
}
