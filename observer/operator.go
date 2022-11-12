package main

import "fmt"

type Observer interface {
	Update(points []Point)
	GetId() string
}

type Operator struct {
	id string
}

func NewOperator(id string) Operator {
	return Operator{id: id}
}

func (operator Operator) Update(points []Point) {
	operator.PrintChanges(points)
}

func (operator Operator) GetId() string {
	return operator.id
}

func (operator Operator) PrintChanges(points []Point) {
	fmt.Printf("Оператор %s: движение в кадре %s \n", operator.id, pointsToString(points))
}
