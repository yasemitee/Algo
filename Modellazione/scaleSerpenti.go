/*
Modellazione:
Nodi: le caselle della scacchiera
Archi: Le mosse migliori (che mi porta più lontano) da una casella ad un altra
é un grafo orientato, connesso, pesato (l'unico caso in cui non è connesso è quando ho per ogni casella una bocca di serpente o una scala)
Per mossa migliore si intende lo spostamento da una casella di partenza a una di arrivo che si avvicini di più alla casella n
(evitando quindi le bocche dei serpenti e prendendo le scale se mi portano più lontano della mossa che tirerei con il dado)

L'obbiettivo è restituire il numero di archi minimo che mi portano dal nodo 1 al nodo n

Implementazione del grafo:
Il grafo è rappresentato in modo implicito tramite la slice "mosseMigliori" dove
-ad ogni posizione della slice è associato un nodo di partenza
-ad ogni valore è associato un arco
Per costruire il grafo con gli archi migliori usiamo la funzione mosseMigliori o mosseMiglioriSenzaSalti
che sceglie l'arco migliore da associare ad un nodo di partenza sfruttando o evitando i salti (se usiamo i salti si scelgono solo i salti che portano più lontano, quindi di base eviterà le bocche dei serpenti)

Algoritmo:
Nella funzione "numeroMosseMinimo" dato in input la posizione di partenza, di arrivo,il grafo e i lanci associati ad ogni arrivo
la funzione stampa i lanci effettuati (che saranno sempre i migliori possibili) e calcola il numero di lanci necessari per arrivare alla casella n
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseInput() (int, int, map[int]int) {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	sc.Scan()
	var r, c, p int //p è la casella di partenza
	fmt.Sscanf(sc.Text(), "%d %d %d", &r, &c, &p)
	n := r * c
	salti := make(map[int]int)
	for sc.Scan() {
		var partenza, arrivo int
		fmt.Sscanf(sc.Text(), "%d %d", &partenza, &arrivo)
		salti[partenza] = arrivo
	}
	return n, p, salti
}

func mosseMigliori(arrivo int, salti map[int]int) ([]int, []int) {
	mosseMigliori := make([]int, arrivo+1)
	lanci := make([]int, arrivo+1)
	l := 0
	for i := 1; i < arrivo; i++ {
		max := i
		for k := 1; k < 7; k++ {
			if i+k <= arrivo {
				if salti[i+k] > max {
					l = k
					max = salti[i+k]
				}
				if i+k > max {
					max = i + k
					l = k
				}
			}
		}
		lanci[i] = l
		mosseMigliori[i] = max
	}
	return mosseMigliori, lanci
}

func isCodaSerpente(casella int, salti map[int]int) bool {
	for _, v := range salti {
		if v == casella {
			return true
		}
	}
	return false
}

func mosseMiglioriSenzaSalti(arrivo int, salti map[int]int) ([]int, []int) {
	mosseMigliori := make([]int, arrivo+1)
	lanci := make([]int, arrivo+1)
	l := 0
	for i := 1; i < arrivo; i++ {
		max := i
		for k := 1; k < 7; k++ {
			if i+k <= arrivo {
				_, ok := salti[i+k]
				if i+k > max && !ok && !isCodaSerpente(i+k, salti) { //vogli evitare anche le code dei serpenti
					max = i + k
					l = k
				}
			}
		}
		lanci[i] = l
		mosseMigliori[i] = max
	}
	fmt.Println(mosseMigliori)
	return mosseMigliori, lanci
}

func numeroMosseMinimo(partenza, arrivo int, mosse []int, lanci []int) {
	count := 0
	i := partenza
	for {
		count++
		fmt.Print(lanci[i], " ")
		if mosse[i] == arrivo {
			break
		}
		i = mosse[i]
		if i > arrivo {
			break
		}
	}
	fmt.Println()
	fmt.Println(count)
}

func main() {
	arrivo, partenza, salti := parseInput()
	mosseMigliori, lanci := mosseMigliori(arrivo, salti)
	//mosseMigliori, lanci := mosseMiglioriSenzaSalti(arrivo, salti)
	numeroMosseMinimo(partenza, arrivo, mosseMigliori, lanci)

}
