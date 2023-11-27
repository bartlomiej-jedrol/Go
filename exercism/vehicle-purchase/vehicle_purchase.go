package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	requireLicense := map[string]bool{
		"car":   true,
		"truck": true,
	}
	return requireLicense[kind]
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	text := " is clearly the better choice."

	if option1 < option2 {
		return option1 + text
	}
	return option2 + text
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var (
		discFirstTier  = 0.8
		discSecondTier = 0.7
		discThirdTier  = 0.5
	)

	switch {
	case age < 3:
		return originalPrice * discFirstTier
	case age < 10:
		return originalPrice * discSecondTier
	default:
		return originalPrice * discThirdTier
	}
}
