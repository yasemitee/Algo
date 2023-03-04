package main

import "fmt"

func stampaSubordinati(tab map[string]string, supervisore string) {
	for tab[supervisore] != "" {
		for sup, sub := range tab {
			if sup == supervisore {
				fmt.Println(sub, " ")
				supervisore = sub
				break
			}
		}
	}
}

func quantiSenzaSubordinati(tab map[string]string) (count int) {
	for _, sub := range tab {
		if cercaSubordinato(tab, sub) {
			count++
		}
	}
	return count
}

func cercaSubordinato(tab map[string]string, supervisore string) bool {
	for sup := range tab {
		if sup == supervisore {
			return false
		}
	}
	return true
}

func supervisore(tab map[string]string, dipendente string) {
	for sup, sub := range tab {
		if sub == dipendente {
			fmt.Println(sup)
		}
	}
}

func stampaImpiegatiSopra(tab map[string]string, dipendente string, capo string) {
	for dipendente != capo {
		for sup, sub := range tab {
			if sub == dipendente {
				fmt.Print(sup, " ")
				dipendente = sup
				fmt.Println(dipendente)
				break
			}
		}
	}
	fmt.Println(capo)
}

func figliCapo(tab map[string]string, capo string) (figli []string) {
	for sup, sub := range tab {
		if sup == capo {
			figli = append(figli, sub)
		}
	}
	return figli
}

func stampaImpiegatiConSupervisore(tab map[string]string) {
	for sup, sub := range tab {
		fmt.Println("Dipendente:", sub, "Supervisore: ", sup)
	}
}

func main() {
	gerarchiaLavoro := make(map[string]string, 0)
	gerarchiaLavoro["Emma"] = "Rosi"
	gerarchiaLavoro["Emma"] = "Yeti"
	gerarchiaLavoro["Yeti"] = "Enrico"
	gerarchiaLavoro["Yeti"] = "Ice"
	gerarchiaLavoro["Rosi"] = "Tamara"
	gerarchiaLavoro["Tamara"] = "Giuseppe"
	gerarchiaLavoro["Giuseppe"] = "Stefano"
	gerarchiaLavoro["Ice"] = "Digio"

	//stampaSubordinati(gerarchiaLavoro, "Rosi")
	//fmt.Println(quantiSenzaSubordinati(gerarchiaLavoro))
	//supervisore(gerarchiaLavoro, "Giuseppe")
	stampaImpiegatiSopra(gerarchiaLavoro, "Giuseppe", "Emma")
	//stampaImpiegatiConSupervisore((gerarchiaLavoro))

}
