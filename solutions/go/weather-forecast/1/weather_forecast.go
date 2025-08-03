// Package weather provides tools to forecast and display the current weather conditio
// weather condition, and includes a function to update and retrieve the forecast.
package weather

// CurrentCondition represents the current weather condition in a city.
var CurrentCondition string

// CurrentLocation represents the current city for which the weather is being forecasted.
var CurrentLocation string

// Forecast sets the CurrentLocation and CurrentCondition variables to the provided cityin the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
