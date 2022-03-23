package models

type BusStopItem struct {
	Vehicle_ID int    `json:"vehicle_id"`
	Arrival    string `json:"arrival_time"`
}

type BusStopOutput struct {
	ID    int           `json:"busstop_id"`
	Name  string        `json:"name"`
	LAT   string        `json:"latitude"`
	LON   string        `json:"longitude"`
	Buses []BusStopItem `json:"buses"`
}

type BusStop struct {
	Forecast []Forecast `json:"forecast"`
	Geometry []Geometry `json:"geometry"`
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	// External_ID  string     `json:"external_id"`
	// Name_EN      string     `json:"name_en"`
	// Name_RU      string     `json:"name_ru"`
	// Nameslug     string     `json:"nameslug"`
	// Resource_URI string     `json:"resource_uri"`
}

type Forecast struct {
	Forecast_Seconds float64 `json:"forecast_seconds"`
	Vehicle_ID       int     `json:"vehicle_id"`
	// Route            Route   `json:"route"`
	// RV_ID            int     `json:"rv_id"`
	// Total_Pass       float64 `json:"total_pass"`
	// Vehicle          string  `json:"vehicle"`
}

type Geometry struct {
	LAT string `json:"lat"`
	LON string `json:"lon"`
	// External_ID int    `json:"external_id"`
	// Seq         int    `json:"seq"`
}

// type Route struct {
// 	ID         int    `json:"id"`
// 	Name       string `json:"name"`
// 	Short_Name string `json:"short_name"`
// }
