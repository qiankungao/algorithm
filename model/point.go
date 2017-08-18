package model

import (
	//	"fmt"
	"math"
)

type Point struct {
	Lable   string
	Entry   []float64
	Visited bool
}

//欧几里得距离
func (point_1 Point) EuclideanDistance(point_2 Point) float64 {
	sum := 0.0
	for i, v := range point_1.Entry {
		sum += math.Pow(v-point_2.Entry[i], 2)
	}

	return math.Sqrt(sum)
}

//坎贝拉距离
func (point_1 Point) CanberraDistance(point_2 Point) (dis float64) {

	for i := range point_1.Entry {
		dis += (math.Abs(point_1.Entry[i]-point_2.Entry[i]) / (math.Abs(point_1.Entry[i]) + math.Abs(point_2.Entry[i])))
	}
	return
}
