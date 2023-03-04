package main

import "fmt"

type Permutazioni struct {
	chiave int
	nome   string
}

func scambia(slice []Permutazioni, j1 int, j2 int) {
	curr := slice[j1]
	slice[j1] = slice[j2]
	slice[j2] = curr
}

// selection sort trovando il massimo ogni volta dell'array non ordinato che va da 0 a prefisso-1
func selectionSortIter(a []Permutazioni) []Permutazioni {
	//n è il prefisso
	n := len(a)
	for i := 0; i < len(a); i++ {
		min := a[0].chiave
		index := 0
		for k := 0; k < n; k++ {
			if a[k].chiave < min {
				min = a[k].chiave
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

func main() {
	pr := []Permutazioni{
		{chiave: 6, nome: "Francesco"},
		{chiave: 1, nome: "Andrea"},
		{chiave: 5, nome: "Elisa"},
		{chiave: 2, nome: "Beatrice"},
		{chiave: 3, nome: "Carlo"},
		{chiave: 4, nome: "Dino"},
		{chiave: 7, nome: "Giorgia"},
		{chiave: 9, nome: "Irene"},
		{chiave: 8, nome: "Henry"},
	}
	fmt.Println(pr)

	fmt.Println("array ordinato: ", selectionSortIter(pr))
}
