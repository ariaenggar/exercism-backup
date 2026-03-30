// Package weather provides weather tools.
package weather

var (
// CurrentCondition for storing current weather condition.
	CurrentCondition string
// CurrentLocation for storing current location.
	CurrentLocation  string
)

// Forecast function returns current location and weather condition string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
