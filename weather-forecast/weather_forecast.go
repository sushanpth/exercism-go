// Package weather forecasts the current weather condition of
// various cites of Goblinocous.
package weather

// CurrentCondition represents current weather condition of the given city.
var CurrentCondition string

// CurrentLocation represents the name of the given city to forecast the weather.
var CurrentLocation string

// Forecast returns a string containing current weather condition of the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
