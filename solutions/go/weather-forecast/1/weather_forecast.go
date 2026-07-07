// Package weather provides a Forecast function
package weather

// CurrentLocation represents the location to forecast
// CurrentCondition represents the condition to forecast
var (
	CurrentCondition string
	CurrentLocation  string
)

// Forecast returns string message provided current location and current condition
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
