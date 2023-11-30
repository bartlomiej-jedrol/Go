package lasagna

const (
	defaultPrepTimePerLayer = 2
	noodlesPortion          = 50
	saucePortion            = 0.2
	defaultPortionNum       = 2
)

// PreparationTime returns an estimated prepartion time of the lasagna.
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = defaultPrepTimePerLayer
	}
	return len(layers) * avgPrepTime
}

// Quantities computes the amount of noodles and saouce needed.
func Quantities(layers []string) (noodlesNeed int, sauceNeed float64) {
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodlesNeed += noodlesPortion
		case "sauce":
			sauceNeed += saucePortion
		}
	}
	return noodlesNeed, sauceNeed
}

// AddSecretIngredient replaces the last ingredient from the list with a secret ingredient.
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// ScaleRecipe returns amounts needed for the desired number of portions.
func ScaleRecipe(quantities []float64, numPortions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		scaledQuantities[i] = quantities[i] * float64(numPortions) / defaultPortionNum
	}
	return scaledQuantities
}
