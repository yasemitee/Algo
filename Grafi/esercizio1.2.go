package main

import (
	"bufio"
	"fmt"
	"os"
)

type grafo struct {
	vertici   []*vertice
	adiacenti map[*vertice][]*vertice
}

type vertice struct {
	nome  string
	eta   int
	hobby []string
}

func graphNew(n int) *grafo {
	var g grafo
	g.vertici = make([]*vertice, n)
	g.adiacenti = make(map[*vertice][]*vertice)
	return &g
}

func newVertex(nome string, eta int, hobby []string) *vertice {
	var v vertice
	v.nome, v.eta, v.hobby = nome, eta, hobby
	return &v
}

func leggiGrafo() *grafo {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Inserire il numero di archi: ")
	var n int
	fmt.Scan(&n)
	//allocazione in memoria un nuovo grafo
	g := graphNew(n)
	//inserimento dei vertici 1 alla volta
	for i := 0; i < n; i++ {
		var nome string
		var eta int
		var hobby []string
		fmt.Println("Inserire i dati del", i+1, "° vertice: ")
		fmt.Print("nome: ")
		fmt.Scan(&nome)
		fmt.Print("eta: ")
		fmt.Scan(&eta)
		fmt.Print("hobby (quando finiti inserire -1): ")
		for scanner.Scan() {
			if scanner.Text() == "-1" {
				break
			}
			hobby = append(hobby, scanner.Text())
		}
		//creo un vertice con i dati e lo inserisco nei vertici del grafo
		g.vertici[i] = newVertex(nome, eta, hobby)
	}

	//creo la lista di adiacenza del grafo map[*vertice][]*vertice
	for i := 0; i < len(g.vertici); i++ {
		var archi []*vertice
		fmt.Println("Inserire i vertici adiacenti di '", g.vertici[i].nome, "' (inserire -1 quando finito):")
		for scanner.Scan() {
			if scanner.Text() == "-1" {
				break
			}
			element := scanner.Text()
			for j := 0; j < len(g.vertici); j++ {
				if g.vertici[j].nome == element {
					archi = append(archi, g.vertici[j])
				}
			}
		}
		g.adiacenti[g.vertici[i]] = archi
	}
	return g
}

func stampaGrafo(g grafo) {
	fmt.Println("Vertici:")
	for i := 0; i < len(g.vertici); i++ {
		fmt.Println(g.vertici[i].nome, g.vertici[i].eta, g.vertici[i].hobby)
	}
	fmt.Println("Lista di adiacenza:")
	for k, v := range g.adiacenti {
		fmt.Print(k.nome, " -> ")
		for _, vertex := range v {
			fmt.Print(vertex.nome, " ")
		}
		fmt.Println("")
	}
}

// mi serve per restituire un vertice dato il nome di un vertice in stringa
func getVertice(g grafo, ver string) (*vertice, bool) {
	for i := 0; i < len(g.vertici); i++ {
		if g.vertici[i].nome == ver {
			return g.vertici[i], true
		}
	}
	return nil, false
}

func stampaHobbyFollower(g grafo, ver string) {
	vertex, err := getVertice(g, ver)
	if err == false {
		fmt.Println("Vertice non trovato nel grafo")
		return
	}
	fmt.Println("Hobby di '", vertex.nome, ":", vertex.hobby)
	for k, v := range g.adiacenti {
		if k == vertex {
			for _, adiacente := range v {
				fmt.Println("Hobby di '", adiacente.nome, ":", adiacente.hobby)
			}
		}
	}
}

func stampaHobbySeguiti(g grafo, ver string) {
	vertex, err := getVertice(g, ver)
	if err == false {
		fmt.Println("Vertice non trovato nel grafo")
		return
	}
	fmt.Println("Hobby di '", vertex.nome, ":", vertex.hobby)
	for k, v := range g.adiacenti {
		for _, adiacente := range v {
			if vertex == adiacente {
				fmt.Println("Hobby di '", k.nome, ":", k.hobby)
			}
		}
	}
}

func main() {
	test := leggiGrafo()
	fmt.Println("Inserire il vertice di cui si vuole sapere gli hobby")
	var s string
	fmt.Scan(&s)
	stampaGrafo(*test)
	//stampaHobbyFollower(*test, s)
	//stampaHobbySeguiti(*test, s)

}
