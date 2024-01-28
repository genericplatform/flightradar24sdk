package main

import (
	"context"
	"fmt"
	"github.com/genericplatform/flightradar24sdk"
)

func main() {
	api := flightradar24sdk.NewAPI(nil)

	resp, err := api.GetFlights(context.Background(), "DAL", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
