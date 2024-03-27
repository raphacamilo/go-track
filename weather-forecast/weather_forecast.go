// Package weather provides tools to view the current weather.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string

// CurrentLocation stores the location of the weather to be viewed.
var CurrentLocation string

// Forecast returns a string containing the weather condition of a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
