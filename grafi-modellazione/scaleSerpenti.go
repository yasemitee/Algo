package main

import (
	"bufio"
	"fmt"
	"os"
)

type grafo struct {
	n     int
	salti []int
}

func nuovoGrafo(n int) *grafo {
	var g grafo
	g.n = n
	g.salti = make([]int, n+1)
	return &g
}

// risolve il problema cercando di andare il più avanti possibile con le scale giuste
func conSalti() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//creazione tavola con nCaselle
	var nCaselle, n, m int
	fmt.Sscanf(sc.Text(), "%d %d", &n, &m)
	nCaselle = n * m
	tavola := nuovoGrafo(nCaselle)
	//mappa con i salti dei serpenti/scale
	salti := make(map[int]int)
	for sc.Scan() {
		var partenza, arrivo int
		fmt.Sscanf(sc.Text(), "%d %d", &partenza, &arrivo)
		salti[partenza] = arrivo
	}
	//metto a posto la lista di adiacenza con i salti
	for i := range tavola.salti {
		for k, v := range salti {
			if k == i {
				tavola.salti[i] = v
			}
		}
	}
	fmt.Println(tavola.salti)
	count := 0
	lanci := 0
	sequenzaDado := make([]int, 0)
	var casellaPartenza int
	fmt.Println("Inserire la casella di partenza:")
	fmt.Scan(&casellaPartenza)
	if casellaPartenza >= nCaselle {
		panic("La casella di partenza non è presente nella tavola")
	}
	for i := casellaPartenza; i < len(tavola.salti); i++ {
		if i == nCaselle {
			sequenzaDado = append(sequenzaDado, count)
			lanci++
			break
		}
		if tavola.salti[i] > i && count <= 6 {
			sequenzaDado = append(sequenzaDado, count)
			count = 0
			lanci++
			i = tavola.salti[i]
			continue
		}
		if count == 6 && tavola.salti[i] < i {
			i = i - 1
			sequenzaDado = append(sequenzaDado, count)
			count = 0
			lanci++
		}
		count++
	}
	fmt.Println(lanci)
	for _, d := range sequenzaDado {
		fmt.Print(d, " ")
	}
	fmt.Println()
}

// risolve il problema evitando le scale e i serpenti
func evitaSalti() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//creazione tavola con nCaselle
	var nCaselle, n, m int
	fmt.Sscanf(sc.Text(), "%d %d", &n, &m)
	nCaselle = n * m
	tavola := nuovoGrafo(nCaselle)
	//mappa con i salti dei serpenti/scale
	salti := make(map[int]int)
	for sc.Scan() {
		var partenza, arrivo int
		fmt.Sscanf(sc.Text(), "%d %d", &partenza, &arrivo)
		salti[partenza] = arrivo
	}
	//metto a posto la lista di adiacenza con i salti
	for i := range tavola.salti {
		for k, v := range salti {
			if k == i {
				tavola.salti[i] = v
			}
		}
	}
	fmt.Println(tavola.salti)
	count := 0
	lanci := 0
	sequenzaDado := make([]int, 0)
	var casellaPartenza int
	fmt.Println("Inserire la casella di partenza:")
	fmt.Scan(&casellaPartenza)
	if casellaPartenza >= nCaselle {
		panic("La casella di partenza non è presente nella tavola")
	}
	for i := casellaPartenza; i < len(tavola.salti); i++ {
		if i == nCaselle {
			fmt.Println(i)
			sequenzaDado = append(sequenzaDado, count)
			lanci++
			break
		}
		if count == 6 {
			fmt.Println(i, "+")
			if tavola.salti[i] != 0 {
				i = i - 1
				fmt.Println(i)
				sequenzaDado = append(sequenzaDado, count-1)
			} else {
				fmt.Println(i)
				sequenzaDado = append(sequenzaDado, count)
			}
			count = 0
			lanci++
		}
		count++
	}
	fmt.Println(lanci)
	for _, d := range sequenzaDado {
		fmt.Print(d, " ")
	}
	fmt.Println()
}

func main() {
	//conSalti()
	//evitaSalti()
}
