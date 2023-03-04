package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type grafo struct {
	vertici   []int
	adiacenti map[int][]int
}

func nuovoGrafo(n int) *grafo {
	var g grafo
	g.vertici = make([]int, 0, n)
	g.adiacenti = make(map[int][]int)
	return &g
}

func parseInput() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	caverna := nuovoGrafo(n * n)
	for sc.Scan() {
		riga := strings.Split(sc.Text(), " ")
		for i, r := range riga {
			rischio, _ := strconv.Atoi(r)
			if i == 0 {
				caverna.vertici = append(caverna.vertici, rischio)
				caverna.adiacenti[rischio] = append(caverna.adiacenti[rischio])
			}
		}
	}
}

func main() {

}
