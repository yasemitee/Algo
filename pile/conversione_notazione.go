package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) push(str string) {
	*s = append(*s, str)
}

func (s *Stack) pop() string {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func converti(espressione string) string {
	elements := strings.Split(espressione, " ")
	var op Stack
	var postFissa string
	for i := 0; i < len(elements); i++ {
		if _, err := strconv.Atoi(elements[i]); err == nil {
			postFissa += elements[i] + " "
		} else {
			if elements[i] != "(" && elements[i] != ")" {
				op.push(elements[i])
			} else if elements[i] == ")" {
				postFissa += op.pop() + " "
			}
		}
	}
	return postFissa
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(converti(input))
}
