package model

import (
	//	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

//欧几里得距离
func Euclideandistance(point, point2 Point) (dis float64) {
	dis = math.Sqrt(math.Pow(point.X-point2.X, 2) + (math.Pow(point.Y-point2.Y, 2)))
	return
}
