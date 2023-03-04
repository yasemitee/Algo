package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type oggetto struct {
	peso   int
	valore int
}

func riempiZaino(pesoZaino int, oggetti []oggetto) int {
	s := make([]int, pesoZaino)
	s[0] = 0
	for i := 1; i < pesoZaino; i++ {
		s[i] = s[i-1] + 
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	pesoZaino := sc.Text()
	var oggetti []oggetto
	for sc.Scan() {
		var o oggetto
		input := strings.Split(sc.Text(), " ")
		o.peso, _ = strconv.Atoi(input[0])
		o.valore, _ = strconv.Atoi(input[1])
		oggetti = append(oggetti, o)
	}
}
