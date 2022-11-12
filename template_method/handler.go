package main

import (
	"fmt"
)

type Handler interface {
	getType(string) string
	getCategory(string) string
	getWeight(int) int
	getDimensions(int) int
	canDeterminePerishable() bool
}

type HandlerIml struct {
	Handler
}

func (handler HandlerIml) getProductPlace(productQR ProductQR) string {
	productType := handler.getType(productQR.pType)
	category := handler.getCategory(productType)
	storeType := handler.getStoreType(category)
	if storeType == notDetermined {
		return notDetermined
	}

	weight := handler.getWeight(productQR.weight)
	dimensions := handler.getDimensions(productQR.dimensions)
	rack := handler.getRack(weight, dimensions)

	return fmt.Sprintf("%s-%s", storeType, rack)
}

func (handler HandlerIml) getRack(weight, dimensions int) string {
	switch {
	case weight < 50 && dimensions < 20:
		return "top"
	case weight > 300 && dimensions > 100:
		return "bottom"
	default:
		return "middle"
	}
}

func (handler HandlerIml) isPerishable(category string) bool {
	perishableProducts := map[string]bool{
		"milky":  true,
		"yogurt": true,
	}

	_, isPerishable := perishableProducts[category]
	return isPerishable
}

func (handler HandlerIml) getStoreType(category string) string {
	if handler.canDeterminePerishable() && handler.isPerishable(category) {
		return "fridge"
	}

	switch category {
	case "food":
		return "foodStore"
	case "clothes":
		return "foodClothes"
	case "technique":
		return "foodTechnique"
	default:
		return notDetermined
	}
}
