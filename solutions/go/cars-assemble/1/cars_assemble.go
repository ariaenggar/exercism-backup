package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate/100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    var x float64 = CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(x)/60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    var _10s int = carsCount / 10
	var remainder int = carsCount % 10

    var cost uint = uint(_10s * 95000) + uint(remainder * 10000)
    
	return cost
}
