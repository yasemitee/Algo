package main

import "fmt"

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	fistHalf := mergeSort(a[:len(a)/2])
	secondHalf := mergeSort(a[len(a)/2:])
	return merge(fistHalf, secondHalf)
}

func merge(a []int, b []int) []int {
	sorted := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			sorted = append(sorted, a[i])
			i++
		} else {
			sorted = append(sorted, b[j])
			j++
		}
	}
	if i < len(a) {
		for ; i < len(a); i++ {
			sorted = append(sorted, a[i])
		}
	} else if j < len(b) {
		for ; j < len(b); j++ {
			sorted = append(sorted, b[j])
		}
	}
	return sorted
}

func main() {
	a := []int{3, 7, 2, 5, 8, 1}
	fmt.Println(mergeSort(a))
}
