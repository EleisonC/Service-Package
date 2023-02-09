package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/EleisonC/Service-Package/configSetup"
	"github.com/EleisonC/Service-Package/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterServiceRoutes(r)
	http.Handle("/", r)
	configs.ConnectDB()
	log.Fatal(http.ListenAndServe(":8080", r))
}