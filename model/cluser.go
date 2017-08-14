package model

//import (
//	"fmt"
//)

type Cluster struct {
	Center   Point
	Elements []Point
}

////Calculate cluster's Center
//func CalculateCenter(cluster Cluster) (center Point) {
//	var sumx, sumy float64

//	for _, point := range cluster.Elements {
//		sumx += point.X
//		sumy += point.Y
//	}
//	center = Point{}
//	num := float64(len(cluster.Elements))
//	cluster.Center.X = sumx / num
//	cluster.Center.Y = sumy / num
//	return cluster.Center
//}
