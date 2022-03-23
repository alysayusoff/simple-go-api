package routes

import (
	"github.com/alysayusoff/bus-service/controllers"
	"github.com/gorilla/mux"
)

var BusServiceRoutes = func(router *mux.Router) {
	router.HandleFunc("/busstop", controllers.GetBusStops)
	router.HandleFunc("/busstop/{busStopId}", controllers.GetBusStopByID)
	router.HandleFunc("/busline", controllers.GetBusLines)
	router.HandleFunc("/busline/{busLineId}", controllers.GetBusLineByID)
}
