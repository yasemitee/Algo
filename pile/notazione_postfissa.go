package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) push(n int) {
	*s = append(*s, n)
}

func (s *Stack) pop() int {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func valuta(espressione string) int {
	elements := strings.Split(espressione, " ")
	var numbers Stack
	for i := 0; i < len(elements); i++ {
		if n, err := strconv.Atoi(elements[i]); err == nil {
			numbers.push(n)
		} else {
			n1, n2 := numbers.pop(), numbers.pop()
			switch elements[i] {
			case "+":
				numbers.push(n1 + n2)
			case "-":
				numbers.push(n2 - n1)
			case "*":
				numbers.push(n1 * n2)
			case "/":
				numbers.push(n2 / n1)
			default:
				fmt.Println("operazione non valida")
			}
		}
	}
	return numbers.pop()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(valuta(input))
}
