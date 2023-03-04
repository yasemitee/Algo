/*
Modellazione:
Usiamo la struttura dati ad albero generico, dove ogni padre è un dipendente e ogni figlio è un subordinato.
ogni dipendente di primo livello è una radice, un insieme di dipendenti di primo livello è una forseta di alberi (che rappresenta la multinazionale)
I compiti da svolgere sono quindi:
-stampare i sotto-alberi di un nodo
-restituire i nodi che sono foglie (senza figli),
-stampare il padre dato un nodo
-stampare tutti i padri e i fratelli dei padri, partendo da un nodo fino alla radice
-per ogni nodo stampare il nodo e il padre. Alla fine stampo anche la radice

Per implementare la struttura viene implementata una struttura che ha come campi:
-tabella dei padri (ottimizza le operazioni di salita).
-una mappa di subordinati, che associa ad ogni nodo una slice di figli, se un dipendente è una foglia avrà associata una slice vuota.

Progettazione:
a)pratire dal nodo (supervisore) dato in input, salvo in una slice di supporto il supervisore e tutti i suoi subordinati tramite la funzione "memorizzaSubordinati"
e poi stampo la slice escludendo il primo elemento (che sarà il supervisore).
la funzione memorizzaSubordinati partendo dal nodo dato in input, in modo ricorsivo scorre la slice di subordinati per ogni dipendente che è subordinato al nodo (input)
b)itero sulla mappa "subordinati" e quando una chiave (dipendente) ha associata una slice vuota (che quindi indica che non ha subordinati) aumento un contatore che verrà restituito
c)dato input un dipendente diverso dal dipendente di primo livello, stampo il suo supervisore
d)In modo ricorsivo stampo i supervisori per ogni supervisore del dipendente fino al dipendente di primo livello
*/

package main

import (
	"fmt"
)

type Azienda = []Team

type Team struct {
	dipendentePrimoLivello string
	supervisore            map[string]string
	subordinati            map[string][]string
	altezza                int
}

func createTeam(nome string) Team {
	var t Team
	t.dipendentePrimoLivello = nome
	t.supervisore = make(map[string]string)
	t.subordinati = make(map[string][]string)
	t.subordinati[nome] = make([]string, 0)
	return t
}

func (t *Team) addNewDipendente(nome string, capo string) {
	t.supervisore[nome] = capo
	t.subordinati[capo] = append(t.subordinati[capo], nome)
	t.subordinati[nome] = []string{}
	t.altezza++
}

func (t Team) memorizzaSubordinati(dipendente string) (sub []string) {
	if len(t.subordinati[dipendente]) == 0 {
		sub = append(sub, dipendente)
		return sub
	}
	sub = append(sub, dipendente)
	for _, dip := range t.subordinati[dipendente] {
		sub = append(sub, t.memorizzaSubordinati(dip)...)
	}
	return sub
}

func (t Team) stampaSubordinati(Team, dipendente string) {
	s := t.memorizzaSubordinati(dipendente)
	fmt.Print("Subordinati: ")
	for i := 1; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()
}

func (t Team) altezzaNodo(dipendente string) (h int) {
	if t.supervisore[dipendente] == "" {
		return h
	}
	h++
	return h + t.altezzaNodo(t.supervisore[dipendente])
}

func (t Team) quantiSenzaSubordinati() int {
	count := 0
	for _, v := range t.subordinati {
		if len(v) < 1 {
			count++
		}
	}
	return count
}

func (t Team) getSupervisore(dipendente string) string {
	return t.supervisore[dipendente]
}

func (t Team) stampaImpiegatiSopra(dipendente string) {
	sup := t.supervisore[dipendente]
	if sup == t.dipendentePrimoLivello {
		fmt.Print(sup, " \n")
		return
	}
	fmt.Print(sup, " ")
	t.stampaImpiegatiSopra(sup)

}

func (t Team) stampaImpiegatiConSupervisore() {
	fmt.Println("Dipendente 1° Livello:", t.dipendentePrimoLivello)
	for k, v := range t.supervisore {
		fmt.Println("dipendente:", k, "supervisore:", v)
	}

}

func stampaAzienda(a Azienda) {
	for _, t := range a {
		fmt.Println("TEAM", t.dipendentePrimoLivello)
		t.stampaImpiegatiConSupervisore()
	}
}

func main() {
	t := createTeam("A")
	t2 := createTeam("G")
	t.addNewDipendente("B", "A")
	t.addNewDipendente("C", "A")
	t.addNewDipendente("D", "A")
	t.addNewDipendente("E", "B")
	t.addNewDipendente("F", "B")
	t.addNewDipendente("I", "F")
	t2.addNewDipendente("H", "G")
	var Algoré Azienda
	Algoré = append(Algoré, t, t2)
	stampaAzienda(Algoré)

}
