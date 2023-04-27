package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) * successRate) / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	perHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(perHour) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tens := carsCount / 10
	rem := carsCount % 10

	return (uint(tens) * 95000) + (uint(rem) * 10000)
}
