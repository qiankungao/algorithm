// LargeDataAlgorithm project main.go
package main

import (
	"BigDataAlgorithm/algorithm"
	"BigDataAlgorithm/common"
	//	"BigDataAlgorithm/model"
	"fmt"
)

func main() {
	path := "data/test1.csv"
	content := common.ReadData(path)
	clus := algorithm.Kmeans(string(content), 2, 0.001)
	for e := range clus {
		fmt.Println("center:", clus[e].Center)
		for _, point := range clus[e].Points {
			fmt.Println("Lable:", point.Lable, " ", point.Entry)
		}

		fmt.Println("##################################")
	}
}

//package main

//import (
//	"BigDataAlgorithm/model"
//	"fmt"
//	"math"
//	"math/rand"
//)

//type Point struct {
//	Entry []float64
//}

//type Centroid struct {
//	center Point
//	Points []Point
//}

//func (p_1 Point) distanceTo(p_2 Point) float64 {
//	sum := float64(0)
//	for e := 0; e < len(p_1.Entry); e++ {
//		sum += math.Pow((p_1.Entry[e] - p_2.Entry[e]), 2)
//	}
//	return math.Sqrt(sum)
//}

//func (c_1 *Centroid) reCenter() float64 {
//	new_Centroid := make([]float64, len(c_1.center.Entry))
//	for _, e := range c_1.Points {
//		for r := 0; r < len(new_Centroid); r++ {
//			new_Centroid[r] += e.Entry[r]
//		}
//	}
//	for r := 0; r < len(new_Centroid); r++ {
//		new_Centroid[r] /= float64(len(c_1.Points))
//	}
//	old_center := c_1.center
//	c_1.center = Point{new_Centroid}
//	return old_center.distanceTo(c_1.center)
//}

//func KMEANS(data []Point, k uint64, DELTA float64) (Centroids []Centroid) {
//	for i := uint64(0); i < k; i++ {
//		Centroids = append(Centroids, Centroid{center: data[rand.Intn(len(data))]})
//	}

//	converged := false
//	for !converged {
//		for i := range data {
//			min_distance := math.MaxFloat64
//			z := 0
//			for v, e := range Centroids {
//				distance := data[i].distanceTo(e.center)
//				if distance < min_distance {
//					min_distance = distance
//					z = v
//				}
//			}
//			Centroids[z].Points = append(Centroids[z].Points, data[i])
//		}
//		max_delta := -math.MaxFloat64
//		for i := range Centroids {
//			movement := Centroids[i].reCenter()
//			if movement > max_delta {
//				max_delta = movement
//			}
//		}
//		if DELTA >= max_delta {
//			converged = true
//			return
//		}
//		for i := range Centroids {
//			Centroids[i].Points = nil
//		}
//	}
//	return Centroids
//}
//func main() {
//	//	data := []Point{}
//	//	data = append(data, Point{[]float64{1.0, 3.0}})
//	//	data = append(data, Point{[]float64{43.0, 7.0}})
//	//	data = append(data, Point{[]float64{2.0, 12.0}})
//	//	data = append(data, Point{[]float64{12.0, 45.0}})
//	//	data = append(data, Point{[]float64{5.0, 2.0}})
//	//	data = append(data, Point{[]float64{12.0, 7.0}})
//	//	data = append(data, Point{[]float64{5.0, 8.0}})
//	//	data = append(data, Point{[]float64{34.0, 65.0}})
//	//	data = append(data, Point{[]float64{1200.0, 700.0}})
//	//	data = append(data, Point{[]float64{500.0, 800.0}})
//	//	data = append(data, Point{[]float64{340.0, 650.0}})

//	//	centroids := KMEANS(data, 3, 0.001)
//	//	for e := range centroids {
//	//		fmt.Println("center:", centroids[e].center)
//	//		fmt.Println("elements:", centroids[e].Points)
//	//		fmt.Println("##################################")
//	//	}
//	//	fmt.Println()
//	p_1 := Point{0, 0}
//	P_2 := Point{3, 4}
//	dis:=P_1.
//}
