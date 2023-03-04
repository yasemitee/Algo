package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getHight(tab map[string]string, node string, h int) int {
	padre := tab[node]
	if padre == "" {
		return h
	}
	h++
	return getHight(tab, padre, h)
}

func main() {

	tabellaPadri := make(map[string]string, 0)

	contents, err := os.Open("test.txt")
	fileScanner := bufio.NewScanner(contents)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		father := strings.Split(fileScanner.Text(), ")")[0]
		node := strings.Split(fileScanner.Text(), ")")[1]
		tabellaPadri[node] = father
	}

	contents.Close()
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	// tabellaPadri["COM"] = ""
	// tabellaPadri["B"] = "COM"
	// tabellaPadri["C"] = "B"
	// tabellaPadri["D"] = "C"
	// tabellaPadri["E"] = "D"
	// tabellaPadri["F"] = "E"
	// tabellaPadri["G"] = "B"
	// tabellaPadri["H"] = "G"
	// tabellaPadri["I"] = "D"
	// tabellaPadri["J"] = "E"
	// tabellaPadri["K"] = "J"
	// tabellaPadri["L"] = "K"

	numeroOrbite := 0
	for node := range tabellaPadri {
		numeroOrbite += getHight(tabellaPadri, node, 0)
	}
	fmt.Println("Il numero totale di orbite:", numeroOrbite)
}
