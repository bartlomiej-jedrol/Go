package lasagna

// PreparationTime returns an estimated prepartion time of the lasagna.
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgPrepTime
}

// Quantities computes the amount of noodles and saouce needed.
func Quantities(layers []string) (noodlesNeed int, sauceNeed float64) {
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodlesNeed += 50
		}
		if layers[i] == "sauce" {
			sauceNeed += 0.2
		}
	}
	return noodlesNeed, sauceNeed
}

// AddSecretIngredient replaces the last ingredient from the list with a secret ingredient.
func AddSecretIngredient(friendList, myList []string) {
	lastItem := friendList[len(friendList)-1]
	myList[len(myList)-1] = lastItem
}

// ScaleRecipe returns amounts needed for the desired number of portions.
func ScaleRecipe(quantities []float64, numPortions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		scaledQuantities[i] = quantities[i] * float64(numPortions) / 2
	}
	return scaledQuantities
}
