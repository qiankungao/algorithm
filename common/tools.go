package common

import (
	//	"fmt"
	//	"io/ioutil"
	//	"log"
	"math/rand"
	//	"path/filepath"
	//	"strconv"
	//	"strings"
	"math"
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
func MaxArray(arr []float64) float64 {
	var max float64

	max = -math.MaxFloat64
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func MinArray(arr []float64) (float64, float64) {
	var min float64
	min = math.MaxFloat64
	for _, v := range arr {
		if v < min {
			min = v
		}
	}

	max := -math.MaxFloat64
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return min, max - min
}
