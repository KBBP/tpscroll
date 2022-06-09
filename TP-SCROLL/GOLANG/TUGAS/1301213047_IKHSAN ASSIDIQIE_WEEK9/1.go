package main

import (
	. "fmt"
	"math"
)

type Point struct {
	x, y int
}

//func PointClass(x int, y int) *Point {
//	return &Point{x: x, y: y}
//} CONSTRUCTOR

func distance(P1, P2 Point) int {
	var cordX, cordY, coordinationXNew, coordinationYNew, total, totalDistance int
	cordX = (P1.x) - (P2.x)
	coordinationXNew = cordX * cordX
	cordY = (P1.y) - (P2.y)
	coordinationYNew = cordY * cordY

	total = coordinationXNew + coordinationYNew
	totalDistance = squareRoot(total)
	return totalDistance
}

func squareRoot(x int) int {
	var roots int
	roots = int(math.Sqrt(float64(x)))
	return roots
}

func main() {
	var P1, P2 Point
	Scanln(&P1.x, &P1.y, &P2.x, &P2.y)
	Println(distance(P1, P2))
}
