/* MODELLAZIONE
NODI: svincoli delle gallerie
ARCHI: gallerie

c'è un arco tra 2 svincoli se esiste una galleria che li colleghi, e il peso su ogni arco indica la luminosità della galleria

il grafo è non orientato, connesso e pesato

si tratta di un problema di esistenza di un cammino, scegliendo sempre e solo l'arco con peso minore

ALGORITMO:  c'è una lista di nodi visitati, un dato che indica il n di archi attraversati, e un dato che indica in che nodo siamo

    se sono in S finito
	se non sono in S: scelgo l'arco con peso minore e mi muovo nel nodo collegato con quell'arco,
	  	si aggiunge 1 alla variabile del numero di archi attraversati
	se il nodo è nella lista dei visitati: non esiste un cammino che rispetti la strategia
	se il nodo non è nella lista si inserisce il nodo nella lista e si ripete l'algoritmo

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type arcoMigliore []int

//pos 0 = vicino migliore
//pos 1 = luminosità

type Grafo = []arcoMigliore //indice = nodo di partenza

func main() {
	var numGallerie int
	var nodiVisitati []int
	var posizione int

	g, h, s := CreaGrafo()

	posizione = h

	for {
		if posizione == s {
			fmt.Println(numGallerie)
			return
		}
		posizione = g[posizione-1][0]
		numGallerie++
		if !(CercaInSlice(nodiVisitati, posizione)) {
			nodiVisitati = append(nodiVisitati, posizione)
		} else {
			fmt.Println(-1)
			return
		}
	}
}

func CreaGrafo() (g Grafo, h int, s int) { // restituisce il grafo con solo le gallerie sceglibili da harmony, il nodo di partenza dell'algoritmo (Harmony) e il nodo d'arrivo(Sarah)
	var num []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		campi := strings.Split(scanner.Text(), " ")
		for _, cont := range campi {
			n, _ := strconv.Atoi(cont)
			num = append(num, n)
		}
		if len(num) == 4 { //solo prima riga input
			g = make(Grafo, num[0])
			h = num[2]
			s = num[3]
			num = nil
			continue
		}

		if (g[num[0]-1] == nil) || (g[num[0]-1][1] > num[2]) {
			g[num[0]-1] = arcoMigliore{num[1], num[2]}
		}
		if (g[num[1]-1] == nil) || (g[num[1]-1][1] > num[2]) {
			g[num[1]-1] = arcoMigliore{num[0], num[2]}
		}

		num = nil
	}
	//fmt.Println(g)
	return g, h, s
}

func CercaInSlice(s []int, n int) bool {
	for _, num := range s {
		if num == n {
			return true
		}
	}
	return false
}

/*
Traccia per l'analisi:

Considerate l’esercizio “Sunnydale” della scheda su “modellazione con grafi” e lo svolgimento proposto in questo programma.
Rileggete le domande nella sezione 1.2 e valutate se il commento fornisce risposte corrette, chiare e complete
alle quattro domande. In caso contrario, come le correggereste/completereste?

Considerate ora il codice:
1. Il grafo che viene manipolato dal programma corrisponde a quanto descritto nel commento?
    1. Se sì, descrivete come è implementato il grafo (è utile fare riferimento ai nomi dei tipi e delle variabili usate, e alle righe di codice rilevanti)
    2. Se no, descrivete quale grafo è effettivamente implementato nel codice descrivete come è implementato (è utile fare riferimento ai nomi dei tipi e delle variabili usate, e alle righe di codice rilevanti)
	--il grado implementato dal codice è diverso dal grafo descritto nei commenti.
	Questo perchè:
		1)il grafo implementato non è connesso (perché non è detto che ci sia un arco che arrivi a sara) e non è pesato (dato che il peso è un valore che viene usato solo come criterio di scelta)
		2)ogni arco è la galleria migliore (non semplicemente una galleria che collega due svincoli)
2. Considerate la definizione del tipo arcoMigliore all riga 31 e verificate come viene usato nelle righe 85-89. Descrivete a parole questo tipo: cosa rappresentano gli elementi della slice? E’ una definizione appropriata? Perché? La modifichereste? Come?
	--f
3. Considerate l’if nelle righe 77-83: è possibile portarlo fuori dal ciclo? Come? E’ opportuno farlo? Perché?
	--è opportuno farlo per separare la parte di allocazione dell'array e numero di volte da iterare (input numero di archi) dalla selezione degli archi, rendendo il codice più leggibile.
	Per farlo basta fare scanner.Scan() e scanner.Text() fuori dal ciclo, e nelle condizioni del ciclo mettere for i < 6 (archi), con i inizialmente = 0
4. Considerate la slice nodiVisitati definita alla riga 40. Cosa rappresenta e come viene usata? E’ una scelta appropriata? In caso negativo, come la modifichereste?
	--serve per tenere traccia dei nodi visitati in modo che verranno visitati tutti i nodi (e la funzione restituisce true) allora sara non è stata trovata.
	Non è una scelta appropriata dato che viene utilizzato spazio aggiuntivo.
	Potremmo iterare sul grafo che essendo non orientato sicuramente da un nodo di partenza arrivera a un nodo di arrivo, quindi ci basta
	usare un contatore che incrementa ad ogni iterazione, quando il contatore è uguale al numero dei nodi (e quindi non abbiamo trovato sara) la funzione restituisce -1
Avete altri commenti o osservazioni relativi allo svolgimento di questo esercizio?

*/
