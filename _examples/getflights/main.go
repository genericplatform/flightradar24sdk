package main

import (
	"context"
	"fmt"
	"github.com/genericplatform/flightradar24sdk"
)

func main() {
	api := flightradar24sdk.NewAPI(flightradar24sdk.WithDebug(true))

	resp, err := api.GetFlights(context.Background(), "DAL", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d flights data received", len(resp.Flights))
}
