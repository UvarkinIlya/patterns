package main

type InaccurateScanner struct {
	HandlerIml
}

func (handler InaccurateScanner) getType(pType string) string {
	return pType
}

func (handler InaccurateScanner) getCategory(productType string) string {
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

func (handler InaccurateScanner) getWeight(weight int) int {
	return weight + 5
}

func (handler InaccurateScanner) getDimensions(dimensions int) int {
	return dimensions - 2
}

func (handler InaccurateScanner) canDeterminePerishable() bool {
	return false
}
