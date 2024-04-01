package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// TemperatureUnit String() method
func (sc TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[sc]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Temperature String() method
func (sc Temperature) String() string {
	return fmt.Sprintf("%v %v", sc.degree, sc.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// SpeedUnit String() method
func (sc SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[sc]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Speed String() method
func (sc Speed) String() string {
	return fmt.Sprintf("%v %v", sc.magnitude, sc.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// MeterologyData String() method
func (sc MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", sc.location, sc.temperature, sc.windDirection, sc.windSpeed, sc.humidity)
}
