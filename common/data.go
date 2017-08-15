package common

import (
	//	"fmt"
	"io/ioutil"
	"log"
	//	"math/rand"

	"path/filepath"
	"strconv"
	"strings"
	//	"time"
)

//从文件中读取原始数据
func ReadData(path string) []byte {
	fp, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

//对原始数据进行初步处理
func DealRawData(content string) ([][]float64, []string) {

	lines := strings.Split(content, "\n")
	rawData := make([][]float64, len(lines))
	rawLable := make([]string, len(lines))

	for index, line := range lines {
		vector := strings.Split(line, ",")
		lable := vector[len(vector)-1]
		vector = vector[:len(vector)-1]
		floatVector := make([]float64, len(vector))
		for jj := range vector {
			floatVector[jj], _ = strconv.ParseFloat(vector[jj], 64)
		}
		rawData[index] = floatVector
		rawLable[index] = lable

	}
	return rawData, rawLable
}

/*
	1.找出每一个属性的最大、最小值，
	2.标准化每一个point的属性
*/
func StandarData(rawData [][]float64) [][]float64 {
	L := len(rawData[0])
	for i := 0; i < L; i++ {
		vector := []float64{}
		for j := 0; j < len(rawData); j++ {
			vector = append(vector, rawData[j][i])
		}
		min, diff := MinArray(vector)
		for k := 0; k < len(rawData); k++ {

			rawData[k][i] = (rawData[k][i] - min) / diff

		}
	}

	return rawData
}
