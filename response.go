package flightradar24sdk

type Flight struct {
	ID               string  // FlightRadar24 ID
	ICAORegistration string  // ICAO 24-bit address The ICAO24 code (sometimes called the Mode S code) is a 24-bit unique number that is assigned to each vehicle or object that can transmit ADS-B messages. It is usually transmitted by aircraft but some airport ground vehicles and multilateration towers also have ICAO24 codes assigned to them.
	Latitude         float64 // Latitude
	Longitude        float64 // Longitude
	Heading          uint8   // Heading in degrees
	Altitude         uint    // Altitude in feet
	Speed            uint    // Speed in knots
	SquawkCode       string  // Squawk code
	RadarID          string  // Radar ID
	ICAOModel        string  // ICAO model type
	Registration     string  // Registration
	Timestamp        int64   // Timestamp
	Origin           string  // Origin airport IATA code
	Destination      string  // Destination airport IATA code
	FlightNumber     string  // Flight number
	IsOnGround       bool    // Is on ground
	RateOfClimb      uint    // Rate of climb in feet per minute
	CallSign         string  // Callsign
	IsGlider         bool    // Is glider
	Company          string  // Company
}

type FlightStatsItem struct {
	ADSB      int
	MLAT      int
	FAA       int
	FLARM     int
	Estimated int
	Satellite int
	Uat       int
	Other     int
}

type FlightStats struct {
	Total   FlightStatsItem `mapstructure:"total"`
	Visible FlightStatsItem `mapstructure:"visible"`
}
