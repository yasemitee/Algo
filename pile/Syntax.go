package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack []string

func (s *Stack) push(str string) {
	*s = append(*s, str)
}

func (s *Stack) pop() string {
	str := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return str
}

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	points := 0
	for sc.Scan() {
		line := sc.Text()
		chunks := make(Stack, 0)
		for _, c := range line {
			ch := string(c)
			if ch != ")" && ch != "]" && ch != "}" && ch != ">" {
				chunks.push(ch)
			} else {
				openChunk := chunks.pop()
				if ch == ")" && openChunk != "(" {
					points += 3
					break
				}
				if ch == "]" && openChunk != "[" {
					points += 57
					break
				}
				if ch == "}" && openChunk != "{" {
					points += 1197
					break
				}
				if ch == ">" && openChunk != "<" {
					points += 25137
					break
				}
			}
		}
	}
	fmt.Println(points)
}
