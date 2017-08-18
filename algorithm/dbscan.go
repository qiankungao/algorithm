package algorithm

import (
	"BigDataAlgorithm/common"
	"BigDataAlgorithm/model"
)

type DenCluster struct {
	points []model.Point
}

//Caaulate the point'3 neighborhood,Excluding point itself
func regionQuery(point model.Point, points []model.Point, eps float64) []model.Point {
	neighborhoods := make([]model.Point, 0)
	for _, neighPot := range points {
		dis := point.EuclideanDistance(neighPot)
		if dis <= eps && dis != 0 {
			neighborhoods = append(neighborhoods, neighPot)
		}
	}
	return neighborhoods
}

func expandDenCluster(neighborhood []model.Point, pots []model.Point, eps float64, minPts int) []model.Point {
	expandPoints := make([]model.Point, 0)
	for _, ps := range neighborhood {
		if !ps.Visited {
			ps.Visited = true
			newneighborhood := regionQuery(ps, pots, eps)
			if len(newneighborhood)+1 >= minPts {
				expandPoints = common.Merge(expandPoints, newneighborhood)
			}
		}
		//如果是这个邻域的噪声点，就加入到该类 即是//对邻域内的所有点，若尚未归类，则将其划入类C
		if ps.Visited
	}
	return expandPoints

}

//
func dbscan(pots []model.Point, eps float64, minPts int) []DenCluster {
	clusters := make([]DenCluster, 0)
	currentClusterN := -1

	for _, point := range pots {
		if !point.Visited {
			point.Visited = true
			neighborhood := regionQuery(p, pots, eps)

			if len(neighborhood)+1 >= minPts {
				currentClusterN += 1
				clusters[currentClusterN].points = append(clusters[currentClusterN].points, point)

			} else {

			}
		}
	}
	return clusters
}
