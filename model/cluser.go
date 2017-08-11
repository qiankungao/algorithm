package model

import (
	"fmt"
)

type Cluster struct {
	Elements []Point
	Center   Point
}

//Calculate cluster's Center
func CalculateCenter(cluster Cluster) (center Point) {
	var sumx, sumy float64

	for _, point := range cluster.Elements {
		sumx += point.X
		sumy += point.Y
	}
	fmt.Println("sumx:", sumx)
	fmt.Println("sumy:", sumy)
	center = Point{}
	num := float64(len(cluster.Elements))
	cluster.Center.X = sumx / num
	cluster.Center.Y = sumy / num
	fmt.Println("center.x", cluster.Center.X)
	fmt.Println("center.Y:", cluster.Center.Y)
	return cluster.Center
}
