package main

import (
	"fmt"
)

const minPointsForNotify = 2

type Point struct {
	x int
	y int
}

func pointsToString(points []Point) string {
	var str string

	for _, point := range points {
		str = fmt.Sprintf("%s (%d, %d),", str, point.x, point.y)
	}

	if len(str) > 0 {
		str = str[:len(str)-1]
	}

	return str
}

type publisher interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyObservers()
}

type Camera struct {
	matrix     [][]int
	movePoints []Point
	observers  []Observer
}

func NewCamera() Camera {
	return Camera{
		matrix:     nil,
		movePoints: make([]Point, 0),
		observers:  make([]Observer, 0),
	}
}

func (camera Camera) NotifyObservers() {
	for _, observer := range camera.observers {
		observer.Update(camera.movePoints)
	}
}

func (camera *Camera) Register(observer Observer) {
	camera.observers = append(camera.observers, observer)
}

func (camera *Camera) Deregister(observer Observer) {
	for i := 0; i < len(camera.observers); i++ {
		if camera.observers[i].GetId() == observer.GetId() {
			camera.observers = append(camera.observers[:i], camera.observers[i+1:]...)
		}
	}
}

func (camera *Camera) getMovePoints() (points []Point) {
	points = make([]Point, 0)

	for i := 0; i < len(camera.matrix); i++ {
		for j := 0; j < len(camera.matrix[i]); j++ {
			if camera.matrix[i][j] == 1 {
				points = append(points, Point{
					x: i,
					y: j,
				})
			}
		}
	}

	return points
}

func (camera *Camera) updateMatrix(newMatrix [][]int) {
	camera.matrix = newMatrix

	camera.movePoints = camera.getMovePoints()

	if len(camera.movePoints) > minPointsForNotify {
		camera.NotifyObservers()
	}
}
