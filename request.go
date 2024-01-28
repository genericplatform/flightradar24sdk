package flightradar24sdk

var baseHeaders = map[string]string{
	"accept-language": "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7",
	"cache-control":   "max-age=0",
	"origin":          "https://www.flightradar24.com",
	"referer":         "https://www.flightradar24.com/",
	"sec-fetch-dest":  "empty",
	"sec-fetch-mode":  "cors",
	"sec-fetch-site":  "same-site",
	"user-agent":      "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
}

var defaultRadarOpts = &RadarOptions{
	FAA:               true,
	Satellite:         true,
	FLARM:             true,
	MLAT:              true,
	ADSB:              true,
	InAir:             true,
	OnGround:          true,
	Inactive:          true,
	Gliders:           true,
	EstimatedPosition: true,
	Stats:             true,
	Limit:             5000,
	MaxAge:            14400,
}

type RadarOptions struct {
	Bounds            []float64 `url:"bounds,int" del:"!"` // [upperLeftLat, upperLefLng, lowerRightLat, lowerRightLng]
	FAA               bool      `url:"faa,int"`            // Use US/Canada data source
	Satellite         bool      `url:"satellite,int"`      //
	FLARM             bool      `url:"flarm,int"`          // Use FLARM data source
	MLAT              bool      `url:"mlat,int"`           // Use MLAT data source
	ADSB              bool      `url:"adsb,int"`           // Use ADS-B data source
	InAir             bool      `url:"air,int"`            // Get in-air aircraft
	OnGround          bool      `url:"onground,int"`       // Get on-ground aircraft
	Inactive          bool      `url:"vehicles,int"`       // Get inactive aircraft on ground
	Gliders           bool      `url:"gliders,int"`        // Get gliders
	EstimatedPosition bool      `url:"gnd,int"`            // Get estimated position
	Stats             bool      `url:"stats,int"`          // Default true
	Limit             int       `url:"limit"`              // Default 5000
	MaxAge            int       `url:"maxage"`             // Default 14400
}
