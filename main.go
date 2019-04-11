// LargeDataAlgorithm project main.go
package main

import (
	"BigDataAlgorithm/model"
	"fmt"
)

func main() {
	point := model.Point{1, 2}
	point2 := model.Point{3, 4.3}
//$$$$$$$$$$$$$$$$$$$$$$$
	point3 := model.Point{5.5, 6}
	point4 := model.Point{4.3, 5.6}
	cluster := model.Cluster{}
	cluster.Elements = []model.Point{point, point2, point3, point4}
	center := model.CalculateCenter(cluster)

	fmt.Println("center:", center)
}
