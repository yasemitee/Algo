package main

import "fmt"

// selection sort trovando il massimo ogni volta dell'array non ordinato che va da 0 a prefisso-1
func selectionSortIter(a []int) []int {
	//n è il prefisso
	n := len(a)
	for i := 0; i < len(a); i++ {
		max := a[0]
		index := 0
		for k := 0; k < n; k++ {
			if a[k] > max {
				max = a[k]
				index = k
			}
		}
		if a[n-1] != a[index] {
			a[n-1], a[index] = a[index], a[n-1]
		}
		n--
	}
	return a
}

func selectionSortRec(a []int, newA []int) []int {
	n := len(a)
	max := a[0]
	for k := 0; k < n; k++ {
		if a[k] > max {
			max = a[k]
		}
	}
	newA[n-1] = max
	if n == 0 || n == 1 {
		return newA
	} else {
		return selectionSortRec(a[0:n-1], newA)
	}
}

func main() {
	numeri := []int{15, 96, 44, 22, 54, 28, 83}
	newA := make([]int, len(numeri))
	fmt.Println(selectionSortIter(numeri))
	fmt.Println(selectionSortRec(numeri, newA))
}
