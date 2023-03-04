package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type listNode struct {
	item int
	next *listNode
}
type linkedList struct {
	head *listNode
}
type grafo struct {
	n         int // numero di vertici
	adiacenti []*linkedList
}

func newNode(val int) *listNode {
	return &listNode{val, nil}
}

func addNewNode(l *linkedList, val int) {
	node := newNode(val)
	node.next = l.head
	l.head = node
}

func printList(l linkedList) {
	p := l.head
	for p != nil {
		fmt.Print(p.item, " ")
		p = p.next
	}
	fmt.Println()
}

func nuovoGrafo(n int) *grafo {
	var g grafo
	g.n = n
	g.adiacenti = make([]*linkedList, n)
	return &g
}

func leggiGrafo() *grafo {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Scrivi prima il numero di vertici e poi le coppie separate da una virgola:")
	scanner.Scan()
	n, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
	coppie := strings.Split(scanner.Text(), " ")
	g := nuovoGrafo(n)

	for v := range g.adiacenti {
		var l linkedList
		for j := 1; j < len(coppie); j++ {
			from, _ := strconv.Atoi(strings.Split(coppie[j], ",")[0])
			to, _ := strconv.Atoi(strings.Split(coppie[j], ",")[1])
			if from == v {
				addNewNode(&l, to)
			}
		}
		g.adiacenti[v] = &l
	}
	fmt.Println("")
	return g
}

func stampaGrafo(g grafo) {
	fmt.Println("Stampa del grafo (lista di adiacenza)")
	for i := 0; i < (g.n); i++ {
		fmt.Print("vertice: ", i, " -> ")
		printList(*g.adiacenti[i])
	}
}

func main() {
	test := leggiGrafo()
	stampaGrafo(*test)
}

// INPUT:
// 4 0,2 0,3 1,2 2,1 3,0
