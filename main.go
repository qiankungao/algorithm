// LargeDataAlgorithm project main.go
package main

import (
	//	"BigDataAlgorithm/algorithm"
	"BigDataAlgorithm/common"
	//	"BigDataAlgorithm/model"
	"fmt"
	//	"math/rand"
	//	"time"
)

func main() {
	//	path := "data/test1.csv"
	//	content := common.ReadData(path)
	//	clus := algorithm.Kmeans(string(content), 2, 0.001)
	//	for e := range clus {
	//		fmt.Println("center:", clus[e].Center)
	//		for _, point := range clus[e].Points {
	//			fmt.Println("Lable:", point.Lable, " ", point.Entry)
	//		}

	//		fmt.Println("##################################")
	//	}
	//	content := common.ReadData(path)
	//	rawData, rawLable := common.DealRawData(string(content))
	//	data := make([]model.Point, len(rawData))
	//	for ii, value := range rawData {
	//		data[ii].Lable = rawLable[ii]
	//		data[ii].Entry = value
	//	}
	//	seeds := algorithm.Seed(data, 2)

	//	fmt.Println("seeds:", seeds)
	rawData := [][]float64{[]float64{2, 1, 102}, []float64{1, 3, 2}}

	fmt.Println(common.StandarData(rawData))

}
