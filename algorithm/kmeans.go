package algorithm

import (
	"BigDataAlgorithm/common"
	"BigDataAlgorithm/model"
	//	"fmt"
	"math"
	"math/rand"
	//	"time"
)

type Cluster struct {
	Center model.Point
	Points []model.Point
}

//初始化k个簇   先采用简单的 即随机选择K 个点
func RandSeed(data []model.Point, k int) (clusters []Cluster) {
	key := common.GenerateRand(k, len(data))
	//	key := []int{0, 1}

	for i := 0; i < k; i++ {
		clusters = append(clusters, Cluster{Center: data[key[i]]})
	}
	return
}

//Initialize the seed with K-Means++
func Seed(data []model.Point, k int) []Cluster {
	clusters := make([]Cluster, k)
	//先随机生成一个种子点
	//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	clusters[0].Center = data[rand.Intn(len(data))]
	d := make([]float64, len(data))
	//求其他的k-1个质点
	for index := 1; index < k; index++ {
		var sum float64
		for kk, p := range data {
			_, dis := near(p, clusters[:index])
			d[kk] = dis
			sum += d[kk]
		}
		target := rand.Float64() * sum
		jj := 0
		for sum = d[0]; sum < target; sum += d[jj] {
			jj++
		}
		clusters[index].Center = data[jj]
	}

	return clusters

}

//calculate the distance from point to the center of cluster
func near(data model.Point, clusters []Cluster) (int, float64) {
	index := 0
	min_dis := math.MaxFloat64
	for ind, v := range clusters {
		dis := data.EuclideanDistance(v.Center)
		if dis < min_dis {
			min_dis = dis
			index = ind
		}
	}
	return index, min_dis
}

//重新计算每一簇的质心
func (c *Cluster) NewCenter() float64 {

	new_center := make([]float64, len(c.Center.Entry))
	for _, p := range c.Points {
		for r := 0; r < len(new_center); r++ {
			new_center[r] += p.Entry[r]
		}
	}

	for i := 0; i < len(new_center); i++ {
		new_center[i] /= float64(len(c.Points))
	}
	old_center := c.Center
	c.Center = model.Point{Entry: new_center}
	return old_center.EuclideanDistance(c.Center)
}

func Kmeans(content string, k int, threshold float64) []Cluster {
	rawData, rawLable := common.DealRawData(content)
	data := make([]model.Point, len(rawData))
	for ii, value := range rawData {
		data[ii].Lable = rawLable[ii]
		data[ii].Entry = value
	}
	seeds := Seed(data, k)
	clus := kmeans(data, seeds, threshold)
	return clus
}

func kmeans(data []model.Point, clusters []Cluster, threshold float64) []Cluster {
	//	clusters := seed(data, k)

	flag := false
	for !flag {
		//计算每一个点与质心的距离，选择最近的质心，并加入它
		for i := range data {
			index, _ := near(data[i], clusters)
			clusters[index].Points = append(clusters[index].Points, data[i])
		}

		//计算每一簇的质心
		max_dis := -math.MaxFloat64
		for index := range clusters {
			diffdis := clusters[index].NewCenter()
			if diffdis > max_dis {
				max_dis = diffdis
			}
		}
		if threshold >= max_dis {
			flag = true
			return clusters
		}
		for index := range clusters {
			clusters[index].Points = nil
		}
	}
	return clusters
}
