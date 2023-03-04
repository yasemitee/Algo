package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// O(n)
func getNumeroAumenti(rilevazioni []int) int {
	count := 0
	for i := 1; i < len(rilevazioni); i++ {
		if rilevazioni[i] > rilevazioni[i-1] {
			count++
		}
	}
	return count
}

// O(n)
func getNumeroAumentiFinestra(rilevazioni []int, m int) int {
	count := 0
	for i := 0; i < len(rilevazioni)-2; i++ {
		if rilevazioni[i] < rilevazioni[i+(m-1)] {
			count++
		}
	}
	return count
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(input[0])
	m, _ := strconv.Atoi(input[1])
	rilevazioni := make([]int, n)
	i := 0
	for i < n {
		sc.Scan()
		rilevazioni[i], _ = strconv.Atoi(sc.Text())
		i++
	}
	fmt.Println(getNumeroAumenti(rilevazioni), getNumeroAumentiFinestra(rilevazioni, m))
}

// in totale O(n)
