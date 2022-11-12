package main

import (
	"fmt"
)

type ProductQR struct {
	pType      string
	weight     int
	dimensions int
}

const notDetermined = "not determined"

func main() {
	apple := ProductQR{
		pType:      "apple",
		weight:     150,
		dimensions: 10,
	}

	notebook := ProductQR{
		pType:      "notebook",
		weight:     3000,
		dimensions: 200,
	}

	milk := ProductQR{
		pType:      "milk",
		weight:     40,
		dimensions: 20,
	}

	handlerWithPerishableScanner := HandlerIml{
		Handler: HandlerWithPerishableScanner{},
	}

	applePlace := handlerWithPerishableScanner.getProductPlace(apple)
	notebookPlace := handlerWithPerishableScanner.getProductPlace(notebook)
	milkPlace := handlerWithPerishableScanner.getProductPlace(milk)

	fmt.Println("PerishableScanner:")
	fmt.Println(applePlace)
	fmt.Println(notebookPlace)
	fmt.Println(milkPlace)

	inaccurateScanner := HandlerIml{
		Handler: InaccurateScanner{},
	}

	applePlaceWrong := inaccurateScanner.getProductPlace(apple)
	notebookPlaceWrong := inaccurateScanner.getProductPlace(notebook)
	milkPlaceWrong := inaccurateScanner.getProductPlace(milk)

	fmt.Println("\nInaccurateScanner:")
	fmt.Println(applePlaceWrong)
	fmt.Println(notebookPlaceWrong)
	fmt.Println(milkPlaceWrong)
}
