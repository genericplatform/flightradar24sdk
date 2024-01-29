package flightradar24sdk

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"net/http"
)

type API struct {
	client *resty.Client
}

func NewAPI(opts ...Option) *API {
	var options apiOptions
	for _, o := range opts {
		o(&options)
	}

	baseClient := http.DefaultClient
	if options.client != nil {
		baseClient = options.client
	}

	r := resty.NewWithClient(baseClient)
	r.SetDebug(options.debugEnabled)
	if options.logger != nil {
		r.SetLogger(options.logger)
	}

	return &API{client: r}
}

func (a *API) Client() *resty.Client {
	return a.client
}

type GetFlightsResponse struct {
	Flights   []Flight    `mapstructure:",remain"`
	Stats     FlightStats `mapstructure:"stats"`
	FullCount float64     `mapstructure:"full_count"`
	Version   float64     `mapstructure:"version"`
}

// GetFlights Get flights:
// - airline: The airline ICAO. Ex: "DAL"
// - radarOpts: Optional radar options.
func (a *API) GetFlights(ctx context.Context, airline string, radarOpts *RadarOptions) (GetFlightsResponse, error) {
	if radarOpts == nil {
		radarOpts = defaultRadarOpts
	}

	v, err := query.Values(radarOpts)
	if err != nil {
		return GetFlightsResponse{}, err
	}

	resp, err := a.client.R().SetContext(ctx).
		SetQueryParamsFromValues(v).
		SetQueryParam("airline", airline).
		SetHeaders(defaultHeaders).
		SetHeader("Accept", "application/json").
		Get(realtimeFlightTrackerDataURL)
	if err != nil {
		return GetFlightsResponse{}, err
	}

	jsonData := make(map[string]any)
	if err := json.Unmarshal(resp.Body(), &jsonData); err != nil {
		return GetFlightsResponse{}, err
	}

	var apiResp struct {
		Stats       FlightStats            `mapstructure:"stats,omitempty"`
		FullCount   float64                `mapstructure:"full_count,omitempty"`
		Version     float64                `mapstructure:"version,omitempty"`
		FlightsData map[string]interface{} `mapstructure:",remain"`
	}
	if err := mapstructure.Decode(jsonData, &apiResp); err != nil {
		return GetFlightsResponse{}, err
	}

	flights := make([]Flight, 0, len(apiResp.FlightsData))
	for k, v := range apiResp.FlightsData {
		data, ok := v.([]interface{})
		if !ok {
			return GetFlightsResponse{}, errors.New("aircraft flights data cast error")
		}
		flights = append(flights, Flight{
			ID:               k,
			ICAORegistration: cast.ToString(data[0]),
			Latitude:         cast.ToFloat64(data[1]),
			Longitude:        cast.ToFloat64(data[2]),
			Heading:          uint8(cast.ToFloat64(data[3])),
			Altitude:         cast.ToUint(data[4]),
			Speed:            uint(cast.ToFloat64(data[5])),
			SquawkCode:       cast.ToString(data[6]),
			RadarID:          cast.ToString(data[7]),
			Registration:     cast.ToString(data[8]),
			ICAOModel:        cast.ToString(data[9]),
			Timestamp:        cast.ToInt64(data[10]),
			Origin:           cast.ToString(data[11]),
			Destination:      cast.ToString(data[12]),
			FlightNumber:     cast.ToString(data[13]),
			IsOnGround:       cast.ToBool(data[14]),
			RateOfClimb:      cast.ToUint(data[15]),
			CallSign:         cast.ToString(data[16]),
			IsGlider:         cast.ToBool(data[17]),
			Company:          cast.ToString(data[18]),
		})
	}

	return GetFlightsResponse{
		Flights:   flights,
		Stats:     apiResp.Stats,
		FullCount: apiResp.FullCount,
		Version:   apiResp.Version,
	}, nil
}
