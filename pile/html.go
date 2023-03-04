package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack []string

func (s *Stack) push(n string) {
	*s = append(*s, n)
}

func (s *Stack) pop() string {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func valuta(espressione string) {
	elements := strings.Split(espressione, " ")
	var tags Stack
	var pos int
	var flag bool
	for i := 0; i < len(elements); i++ {
		ch := []rune(elements[i])
		if len(ch) == 3 {
			tags.push(elements[i])
		} else if ch[1] == '/' {
			t := tags.pop()
			ch2 := []rune(t)
			if ch2[1] != ch[2] {
				fmt.Println("errore in pos", pos)
			} else {
				flag = true
			}
		}
	}
	if flag == true {
		if len(tags) != 0 {
			fmt.Print("sono rimasti aperti i seguenti tag: ")
			for i := 0; i < len(tags); i++ {
				fmt.Print(tags[i])
			}
			fmt.Println()
		} else {
			fmt.Println("Documento ben formato")
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	valuta(input)
}
