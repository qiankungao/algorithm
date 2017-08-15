package common

import (
	//	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//随机生成k个不同的随机数
func GenerateRand(k, rang int) (rint []int) {

	if k > rang {
		return nil
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		if len(rint) == k {
			break
		}
		num := r.Intn(rang)
		flag := false
		for _, v := range rint {
			if num == v {
				flag = true
				break
			}
		}
		if flag == false {
			rint = append(rint, num)
		}
	}
	return
}

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
