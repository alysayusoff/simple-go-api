package models

type BusLineItem struct {
	Vehicle_ID int    `json:"vehicle_id"`
	LAT        string `json:"latitude"`
	LON        string `json:"longitude"`
}

type BusLineOutput struct {
	RV_ID     int           `json:"route_id"`
	Name      string        `json:"name"`
	RouteName string        `json:"routename"`
	Buses     []BusLineItem `json:"buses"`
}

type BusLine struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	RouteName string    `json:"routename"`
	Vehicle   []Vehicle `json:"vehicles"`
	// External_ID  string    `json:"external_id"`
	// Name_EN      string    `json:"name_en"`
	// Name_RU      string    `json:"name_ru"`
	// Nameslug     string    `json:"nameslug"`
	// Resource_URI string    `json:"resource_uri"`
	// Via          string    `json:"via"`
}

type Vehicle struct {
	Projection Projection `json:"projection"`
	Vehicle_ID int        `json:"vehicle_id"`
	// Bearing           int        `json:"bearing"`
	// Device_TS         string     `json:"device_ts"`
	// Enterprise        Enterprise `json:"enterprise"`
	// LAT               string     `json:"lat"`
	// LON               string     `json:"lon"`
	// Park              Park       `json:"park"`
	// Position          Position   `json:"position"`
	// Registration_Code string     `json:"registration_code"`
	// RouteVariant_ID   int        `json:"routevariant_id"`
	// Speed             string     `json:"speed"`
	// Stats             Stats      `json:"stats"`
	// TS                string     `json:"ts"`
}

type Projection struct {
	LAT string `json:"lat"`
	LON string `json:"lon"`
	// Edge_Distance      string `json:"edge_distance"`
	// Edge_ID            int    `json:"edge_id"`
	// Edge_Projection    string `json:"edge_projection"`
	// Edge_Start_Node_ID string `json:"edge_start_node_id"`
	// Edge_Stop_Node_ID  string `json:"edge_stop_node_id"`
	// Orig_LAT           string `json:"orig_lat"`
	// Orig_LON           string `json:"orig_lon"`
	// RouteVariant_ID    int    `json:"routevariant_id"`
	// TS                 int64  `json:"ts"`
}

// type Enterprise struct {
// 	Enterprise_ID   int    `json:"enterprise_id"`
// 	Enterprise_Name string `json:"enterprise_name"`
// }

// type Park struct {
// 	Park_ID   int    `json:"park_id"`
// 	Park_Name string `json:"park_name"`
// }

// type Position struct {
// 	Bearing   int    `json:"bearing"`
// 	Device_TS int64  `json:"device_ts"`
// 	LAT       string `json:"lat"`
// 	LON       string `json:"lon"`
// 	Speed     int    `json:"speed"`
// 	TS        int64  `json:"ts"`
// }

// type Stats struct {
// 	Avg_Speed     string `json:"avg_speed"`
// 	Bearing       int    `json:"bearing"`
// 	Cumm_Speed_10 string `json:"cumm_speed_10"`
// 	Cumm_Speed_2  string `json:"cumm_speed_2"`
// 	Device_TS     int    `json:"device_ts"`
// 	LAT           string `json:"lat"`
// 	LON           string `json:"lon"`
// 	Speed         int    `json:"speed"`
// 	TS            int64  `json:"ts"`
// }
