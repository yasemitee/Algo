package main

import "fmt"

func hanoi(n, from, temp, to int) {
	if n < 2 {
		fmt.Println(from, "->", to)
	}
	hanoi(n-1, from, to, temp)
	fmt.Println(from, "->", to)
	hanoi(n-1, temp, from, to)
}

func main() {
	hanoi(3, 0, 1, 2)
}
