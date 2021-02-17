package vatsim

// APIResponse is the response received from vatsim datafeed v3
type APIResponse struct {
	General     general
	Pilots      []pilot
	Controllers []controller
}

type general struct {
	Version          int    `json:"version"`
	Reload           int    `json:"reload"`
	Update           string `json:"update"`
	UpdateTimestamp  string `json:"update_timestamp"`
	ConnectedClients int    `json:"connected_clients"`
	UniqueUsers      int    `json:"unique_users"`
}

type pilot struct {
	CID         int        `json:"cid"`
	Name        string     `json:"name"`
	Callsign    string     `json:"callsign"`
	Server      string     `json:"server"`
	PilotRating int        `json:"pilot_rating"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	Altitude    int        `json:"altitude"`
	Groundspeed int        `json:"groundspeed"`
	Transponder string     `json:"transponder"`
	Heading     int        `json:"heading"`
	QnhIHg      float64    `json:"qnh_i_hg"`
	QnhMb       int        `json:"qnh_mb"`
	Flightplan  flightPlan `json:"flight_plan"`
	LogonTime   string     `json:"logon_time"`
	LastUpdated string     `json:"last_updated"`
}

type flightPlan struct {
	FlightRules string `json:"flight_rules"`
	Aircraft    string `json:"aircraft"`
	Departure   string `json:"departure"`
	Arrival     string `json:"arrival"`
	Alternate   string `json:"alternate"`
	CruiseTas   string `json:"cruise_tas"`
	Altitude    string `json:"altitude"`
	Deptime     string `json:"deptime"`
	EnrouteTime string `json:"enroute_time"`
	FuelTime    string `json:"fuel_time"`
	Remarks     string `json:"remarks"`
	Route       string `json:"route"`
}

type controller struct {
	CID         int    `json:"cid"`
	Name        string `json:"name"`
	Callsign    string `json:"callsign"`
	Frequency   string `json:"frequency"`
	Facility    int    `json:"facility"`
	Rating      int    `json:"rating"`
	Server      string `json:"server"`
	VisualRange int    `json:"visual_range"`
	LastUpdated string `json:"last_updated"`
	LogonTime   string `json:"logon_time"`
}
