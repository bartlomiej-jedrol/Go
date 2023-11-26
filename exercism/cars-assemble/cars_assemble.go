package cars

const MinutesInHour = 60

var (
	productionRate int     = 1567
	successRate    float64 = 90
	carsBundleSize int     = 10
	standardCost   int     = 10000
	reducedCost    int     = 9500
	carsCount      int     = 11
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / MinutesInHour
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var div = carsCount / carsBundleSize
	var rem = carsCount % carsBundleSize

	return uint(div*carsBundleSize*reducedCost + rem*standardCost)
}
