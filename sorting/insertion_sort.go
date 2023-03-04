package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// insertion sort passando un numero
func insertionSort(num []int, n int) []int {
	for i := 1; i < len(num); i++ {
		n := num[i]
		j := i - 1
		for j >= 0 && num[j] > n {
			num[j+1] = num[j]
			j--
		}
		num[j+1] = n
	}
	return num
}

func main() {
	num := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	//leggo l'input, se legge 0 esce dal for altrimenti aggiunge l'elemento nella posizione giusta
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		num = append(num, n)
		insertionSort(num, n)
		fmt.Println(num)
	}
	fmt.Println(num)
}
