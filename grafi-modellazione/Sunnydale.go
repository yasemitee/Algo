/*
MODELLAZIONE:
 1. I nodi rappresentano gli svincoli
 2. Gli archi rappresentano le gallerie che sceglierebbe Harmony
 3. Il Grafo è orientato, non connesso e pesato
 4. stabilire l'esistenza di un cammino in un grafo.
    In particolare partendo da un nodo specifico bisogna stabilire se si possa raggiungere un determinato nodo con la seguente regola:
    ad ogni passo si sceglie di attraversare l'arco con il peso minimo,

ALGORITMO:

	Struttura del grafo:
	definiamo una struct galleria con un campo che identifica un nodo di arrivo e un campo che identifica la luminosita
	dichiariamo una slice di grandezza n+1, dove ogni indice è un nodo di partenza e ogni valore è la galleria meno luminosa d'arrivo.

	Nella parte di parse dell'input l'algoritmo aggiorna la slice associando ad ogni partenza la galleria di luminosità minore.

	Adesso per verificare l'esistenza di un cammino iteriamo fino a quando il numero di iterazioni (quindi spostamenti effettuati nel grafo) è minore del numero dei nodi.
	Se trovo un nodo di arrivo che corrisponde a sara allora il cammino esiste viene restituito il numero di passi (gallerie attraversate),
	se invece finendo il numero di passi non ho trovato sara allora il cammino non esiste, quindi restituisco -1
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type galleria struct {
	verticeArrivo int
	luminosita    int
}

func parseInput() (gallerie []*galleria, h int, s int) {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(input[0])
	h, _ = strconv.Atoi(input[2])
	s, _ = strconv.Atoi(input[3])
	gallerie = make([]*galleria, n+1)
	for scanner.Scan() {
		arco := strings.Split(scanner.Text(), " ")
		partenza, _ := strconv.Atoi(arco[0])
		arrivo, _ := strconv.Atoi(arco[1])
		peso, _ := strconv.Atoi(arco[2])
		nuovaGalleria1 := &galleria{arrivo, peso}
		nuovaGalleria2 := &galleria{partenza, peso}

		if gallerie[partenza] == nil {
			gallerie[partenza] = nuovaGalleria1
		} else if gallerie[partenza].luminosita > nuovaGalleria1.luminosita {
			gallerie[partenza] = nuovaGalleria1
		}

		if gallerie[arrivo] == nil {
			gallerie[arrivo] = nuovaGalleria2
		} else if gallerie[arrivo].luminosita > nuovaGalleria2.luminosita {
			gallerie[arrivo] = nuovaGalleria2
		}
	}
	return gallerie, h, s
}

func trovaCammino(gallerie []*galleria, h int, s int) {
	count := 1
	i := h
	for {
		if gallerie[i].verticeArrivo == s {
			fmt.Println(count)
			break
		}
		i = gallerie[i].verticeArrivo
		count++
		if count > len(gallerie)-1 {
			fmt.Println(-1)
			break
		}
	}
}

func main() {
	a, h, s := parseInput()
	trovaCammino(a, h, s)
}
