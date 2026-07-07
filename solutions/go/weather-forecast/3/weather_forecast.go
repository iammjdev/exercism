// Package weather provides a Forecast function.
package weather

var (
	// CurrentCondition represents the condition to forecast.
	CurrentCondition string
	// CurrentLocation represents the location to forecast.
	CurrentLocation string
)

// Forecast returns string message provided current location and current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
