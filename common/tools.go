package common

import (
	//	"fmt"
	"math/rand"
	"time"
)

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
