/*
Modellazione:
a)
Nodi: Incroci
Archi: tratti di strada
è un grafo non orientato, connesso, pesato
b)
Possiamo trovare il minimo albero ricoprente del grafo usando l'algoritmo di Kruskal,
gli archi appartenenti al minimo albero ricoprente saranno le tubazioni principali
tutti gli altri archi saranno tratti di strada su cui passeranno le tubazioni secondarie
c)
Il grafo può essere rappresentato come un insieme di numeri interi da 1 a n che rappresentano gli incroci,
avremo poi un insieme di strade che sono gli archi.
Una strada è una struttura formata da {partenza,arrivo e lunghezza}
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Città struct {
	incroci []int
	strade  []strada
}

type strada struct {
	partenza  int
	arrivo    int
	lunghezza int
}

func newCittà(n int) *Città {
	var c Città
	c.incroci = make([]int, n)
	c.strade = make([]strada, n)
	return &c
}

func parseInput() *Città {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	città := newCittà(n)
	for sc.Scan() {
		var partenza, arrivo, lunghezza int
		fmt.Sscanf(sc.Text(), "%d %d %d", &partenza, &arrivo, &lunghezza)
		stradaAndata := &strada{partenza, arrivo, lunghezza}
		stradaRitorno := &strada{arrivo, partenza, lunghezza}
		città.strade = append(città.strade, *stradaAndata)
		città.strade = append(città.strade, *stradaRitorno)
	}
	return città
}

func getRitorno(s strada) *strada {
	return &strada{s.arrivo, s.partenza, s.lunghezza}
}

func Kruskal(c Città) ([]strada, []strada) {
	sort.Slice(c.strade, func(i, j int) bool {
		return c.strade[i].lunghezza < c.strade[j].lunghezza
	})
	primarie := make([]strada, 0)
	secondarie := make([]strada, 0)
	visitati := make(map[int]bool) //partizioni
	andataRitorno := make(map[strada]bool)
	for _, strada := range c.strade {
		ritornoAux := getRitorno(strada)
		if !visitati[strada.partenza] || !visitati[strada.arrivo] {
			primarie = append(primarie, strada)
			//union
			visitati[strada.partenza] = true
			visitati[strada.arrivo] = true
		} else if !andataRitorno[strada] {
			secondarie = append(secondarie, strada)

		}
		andataRitorno[*ritornoAux] = true
	}
	return primarie, secondarie
}

func main() {
	c := parseInput()
	p, s := Kruskal(*c)
	fmt.Println("Principali:", p)
	fmt.Println("Secondarie:", s)
}
