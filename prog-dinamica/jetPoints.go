package main

import (
	"fmt"
)

func max(j int64, i int64) (int64, bool) {
	if j > i {
		return j, true
	}
	return i, false
}

func contains(s []int64, e int) bool {
	for _, a := range s {
		if a == int64(e) {
			return true
		}
	}
	return false
}

func main() {
	jetPoints := []int{20000, 22000, 24000, 26000, 28000, 30000, 32000, 34000, 36000, 38000, 40000}
	swindleValues := []int{100, 107, 124, 133, 139, 155, 172, 178, 184, 190, 195}
	var points int
	fmt.Scan(&points)
	capacity := points
	nOggetti := len(jetPoints)

	values := make([][]int64, nOggetti+1)
	for i := range values {
		values[i] = make([]int64, capacity+1)
	}

	keep := make([][]int, nOggetti+1)
	for i := range keep {
		keep[i] = make([]int, capacity+1)
	}

	for i := 0; i < capacity+1; i++ {
		values[0][i] = 0
		keep[0][i] = 0
	}

	for i := 0; i < nOggetti+1; i++ {
		values[i][0] = 0
		keep[i][0] = 0
	}

	for i := 1; i <= nOggetti; i++ {
		for c := 1; c <= capacity; c++ {

			itemFits := (jetPoints[i-1] <= c)
			if !itemFits {
				continue
			}

			maxValueAtThisCapacity := int64(jetPoints[i-1]) + values[i-1][c-jetPoints[i-1]]
			previousValueAtThisCapacity := values[i-1][c]

			if itemFits && (maxValueAtThisCapacity > previousValueAtThisCapacity) {
				values[i][c] = maxValueAtThisCapacity
				keep[i][c] = 1
			} else {
				values[i][c] = previousValueAtThisCapacity
				keep[i][c] = 0
			}
		}
	}
	n := nOggetti
	c := capacity
	var indices []int64
	for n > 0 {
		if keep[n][c] == 1 {
			indices = append(indices, int64(n-1))
			c -= jetPoints[n-1]
		}
		n--
	}
	sum := 0
	for i := 0; i < len(swindleValues); i++ {
		if contains(indices, i) {
			sum += swindleValues[i]
		}
	}
	fmt.Println(sum)
}
