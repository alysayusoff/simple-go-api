// Creating a RESTful API With Golang - tutorialedge.net/golang/creating-restful-api-with-golang/
// Consuming A RESTful API With Go    - tutorialedge.net/golang/consuming-restful-api-with-go/
// Rate limiting -  github.com/striversity/gotr/blob/master/miscellaneous/ep004-rate-limiter/ex01/main.go

package main

import (
	"log"
	"net/http"

	"github.com/alysayusoff/bus-service/routes"
	"github.com/gorilla/mux"
)

func main() {
	// create a new mux router
	// gorilla/mux router will allow us to retrieve path and query parameters
	router := mux.NewRouter().StrictSlash(true)
	// pass the router into routes.BusServiceRoutes function
	routes.BusServiceRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
