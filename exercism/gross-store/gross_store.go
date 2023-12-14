package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if v, ok := units[unit]; !ok {
		return false
	} else {
		bill[item] += v
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// Validations
	currentQty, ok := bill[item]
	if !ok {
		return false
	}

	reducedQty, ok := units[unit]
	if !ok {
		return false
	}

	// Actual business logic
	targetQty := currentQty - reducedQty
	if targetQty < 0 {
		return false
	} else if targetQty == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] = targetQty
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	currentQty, ok := bill[item]
	if !ok {
		return 0, false
	}
	return currentQty, true
}
