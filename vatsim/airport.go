package vatsim

import (
	"strings"
)

// Airport groups fields related to a specific airport by ICAO
type Airport struct {
	Icao        string
	Controllers []controller
	Arrivals    []string
	Departures  []string
}

// FindAirport creates an Airport from APIResponse
func FindAirport(icao string, data APIResponse) (foundAirport Airport) {
	foundAirport = Airport{
		Icao: strings.ToUpper(icao),
	}

	for i := 0; i < len(data.Controllers); i++ {
		if len(data.Controllers[i].Callsign) < 4 {
			continue
		}

		callsign := strings.ToLower(data.Controllers[i].Callsign[0:4])

		if callsign == icao {
			foundAirport.Controllers = append(foundAirport.Controllers, data.Controllers[i])
		}
	}

	for i := 0; i < len(data.Pilots); i++ {
		departure := strings.ToLower(data.Pilots[i].Flightplan.Departure)
		arrival := strings.ToLower(data.Pilots[i].Flightplan.Arrival)

		if departure == icao {
			foundAirport.Departures = append(foundAirport.Departures, data.Pilots[i].Callsign)
		}

		if arrival == icao {
			foundAirport.Arrivals = append(foundAirport.Arrivals, data.Pilots[i].Callsign)
		}
	}

	return
}
