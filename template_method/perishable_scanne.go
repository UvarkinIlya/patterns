package main

type HandlerWithPerishableScanner struct {
	HandlerIml
}

func (handler HandlerWithPerishableScanner) getType(pType string) string {
	return pType
}

func (handler HandlerWithPerishableScanner) getCategory(productType string) string {
	categories := map[string]string{
		"apple":    "food",
		"T-shirt":  "clothes",
		"milk":     "food",
		"notebook": "technique",
	}

	category, found := categories[productType]
	if !found {
		return notDetermined
	}

	return category
}

func (handler HandlerWithPerishableScanner) getWeight(weight int) int {
	return weight
}

func (handler HandlerWithPerishableScanner) getDimensions(dimensions int) int {
	return dimensions
}

func (handler HandlerWithPerishableScanner) canDeterminePerishable() bool {
	return true
}
