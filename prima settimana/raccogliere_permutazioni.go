package main

import (
	"fmt"
)

func main() {
	var n = []int{4, 5, 1, 3, 6, 2}
	var count int
	rc := 1
	for rc != len(n) {
		for i := 0; i < len(n); i++ {
			if n[i] == rc {
				rc++
			}
		}
		count++
	}
	fmt.Println(count)
}
