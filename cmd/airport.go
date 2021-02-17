package cmd

import (
	"fmt"
	"math"
	"os"

	"github.com/doylemark/vatsim-cli/validate"
	"github.com/doylemark/vatsim-cli/vatsim"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var airportCmd = &cobra.Command{
	Use:     "airport",
	Short:   "vatsim-cli airport <icao>: Displays information about a specified airport",
	Example: "vatsim-cli airport eidw",
	Run: func(cmd *cobra.Command, args []string) {
		err := validate.Args(1, args, validateAirport)

		if err != nil {
			fmt.Println(err)
			return
		}

		handleAirport(args)
	},
}

func init() {
	rootCmd.AddCommand(airportCmd)
}

func handleAirport(args []string) {
	data := vatsim.GetDatafeed()

	icao := args[0]

	airport := vatsim.FindAirport(icao, data)

	controllerCallsigns := []string{}

	for i := 0; i < len(airport.Controllers); i++ {
		controllerCallsigns = append(controllerCallsigns, airport.Controllers[i].Callsign)
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Controllers", "Arrivals", "Departures"})

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgWhiteColor},
	)

	largestMovement := math.Max(float64(len(airport.Arrivals)), float64(len(airport.Departures)))
	largestItem := int(math.Max(largestMovement, float64(len(airport.Controllers))))

	for i := 0; i < largestItem; i++ {
		row := []string{"", "", ""}

		if len(airport.Controllers) > i {
			row[0] = airport.Controllers[i].Callsign
		}

		if len(airport.Arrivals) > i {
			row[1] = airport.Arrivals[i]
		}

		if len(airport.Departures) > i {
			row[2] = airport.Departures[i]
		}

		table.Append(row)
	}

	table.Render()
}

func validateAirport(args []string) (err error) {
	if len(args[0]) != 4 {
		err = fmt.Errorf("ICAO code must be 4 characters long")
	}

	return
}
