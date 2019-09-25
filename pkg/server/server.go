package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	routes "github.com/blockchain-abstraction-middleware/project-template/pkg/routes"
	utils "github.com/blockchain-abstraction-middleware/project-template/pkg/utils"
)

// Serve the service
func Serve() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	fmt.Println("Environment:", viper.Get("environment"))

	r := mux.NewRouter()

	r.HandleFunc("/", utils.Chain(routes.Status, routes.Hello(), routes.Logging()))
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
