package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alysayusoff/bus-service/models"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

var BusStop models.BusStop
var BusLine models.BusLine

func wait() {
	fmt.Printf("waiting... ")
	r := rate.Every(time.Second)   // rate of one second
	lim := rate.NewLimiter(r, 100) // limit API calls to 100 calls per second
	lim.Wait(context.Background()) // wait until able to call again
}

func GetBusStops(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bus Stop Ids to choose from:\n")

	busStopIds := []string{
		"378204", "383050", "378202", "383049", "382998", "378237", "378233", "378230",
		"378229", "378228", "378227", "382995", "378224", "378226", "383010", "383009",
		"383006", "383004", "378234", "383003", "378222", "383048", "378203", "382999",
		"378225", "383014", "383013", "383011", "377906", "383018", "383015", "378207",
	}

	for i := 0; i < len(busStopIds); i++ {
		fmt.Fprintln(w, busStopIds[i])
	}
}

func GetBusLines(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bus Line Ids to choose from:\n")

	busLineIds := []string{
		"44478", "44479", "44480", "44481",
	}

	for i := 0; i < len(busLineIds); i++ {
		fmt.Fprintln(w, busLineIds[i])
	}
}

func GetBusStopByID(w http.ResponseWriter, r *http.Request) {
	// print the time request was made
	wait()
	fmt.Printf("request to API api/platformbusarrival/ made at     : %v\n", time.Now())

	// get busStopId
	vars := mux.Vars(r)
	key := vars["busStopId"]

	// query API endpoint and map results into response, err
	url := "https://baseride.com/routes/api/platformbusarrival/" + key
	response, err := http.Get(url)
	// check if there are eny errors
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// read in data from the incoming byte stream
	responseData, err := ioutil.ReadAll(response.Body)
	// check if there are eny errors
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal the returned JSON string into BusStop
	json.Unmarshal(responseData, &BusStop)

	// insert relevant values into busStopOutput
	busStopOutput := models.BusStopOutput{
		ID:   BusStop.ID,
		Name: BusStop.Name,
		LAT:  BusStop.Geometry[0].LAT,
		LON:  BusStop.Geometry[0].LON,
	}
	// for each forecast available
	for i := 0; i < len(BusStop.Forecast); i++ {
		// calculate the forecast in minutes and seconds and convert into a string for arrival time
		forecast := BusStop.Forecast[i].Forecast_Seconds
		minutes := forecast / 60
		seconds := forecast - (float64(int(minutes)) * 60.0)
		arrival := strconv.Itoa(int(minutes)) + " min " + strconv.Itoa(int(seconds)) + " secs"
		// insert relevant values to show bus numberand arrival time to busStopItem
		busStopItem := models.BusStopItem{
			Vehicle_ID: BusStop.Forecast[i].Vehicle_ID,
			Arrival:    arrival,
		}
		// append busStopItem into busStopOutput.Buses
		busStopOutput.Buses = append(busStopOutput.Buses, busStopItem)
	}

	// apply indent to format the output
	b, err := json.MarshalIndent(busStopOutput, "", "  ")
	if err != nil {
		fmt.Print(err.Error())
	}
	// print output
	fmt.Fprintln(w, string(b))
}

func GetBusLineByID(w http.ResponseWriter, r *http.Request) {
	// print the time request was made
	wait()
	fmt.Printf("request to API apigeo/routevariantvehicle/ made at : %v\n", time.Now())

	// get busLineId
	vars := mux.Vars(r)
	key := vars["busLineId"]

	// query API endpoint and map results into response, err
	url := "https://baseride.com/routes/apigeo/routevariantvehicle/" + key
	response, err := http.Get(url)
	// check if there are eny errors
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// read in data from the incoming byte stream
	responseData, err := ioutil.ReadAll(response.Body)
	// check if there are eny errors
	if err != nil {
		log.Fatal(err)
	}
	// unmarshal the returned JSON string into BusStop
	json.Unmarshal(responseData, &BusLine)

	// insert relevant values into busLineOutput
	busLineOutput := models.BusLineOutput{
		RV_ID:     BusLine.ID,
		Name:      BusLine.Name,
		RouteName: BusLine.RouteName,
	}
	// for each bus available
	for i := 0; i < len(BusLine.Vehicle); i++ {
		// insert relevant values to show bus number, and location (lat, lon) into busLineItem
		busLineItem := models.BusLineItem{
			Vehicle_ID: BusLine.Vehicle[i].Vehicle_ID,
			LAT:        BusLine.Vehicle[i].Projection.LAT,
			LON:        BusLine.Vehicle[i].Projection.LON,
		}
		// append busLineItem into busLineOutput.Buses
		busLineOutput.Buses = append(busLineOutput.Buses, busLineItem)
	}

	// apply indent to format the output
	b, err := json.MarshalIndent(busLineOutput, "", "  ")
	if err != nil {
		fmt.Print(err.Error())
	}
	// print output
	fmt.Fprintln(w, string(b))
}
