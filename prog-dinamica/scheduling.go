package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type intervallo struct {
	Inizio int
	Fine   int
	Peso   int
}

func max(i int, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}

func sommaIntervalli(intervalli []intervallo) int {
	n := len(intervalli)
	m := make([]int, n)
	m[0] = intervalli[0].Peso

	for i := 1; i < n; i++ {
		m[i] = intervalli[i].Peso
		if intervalli[i-1].Fine <= intervalli[i].Inizio {
			m[i] = max(m[i-1], m[i]+intervalli[i].Peso)
		} else {
			for j := i; j >= 0; j-- {
				if intervalli[j].Fine <= intervalli[i].Inizio {
					m[i] = max(m[i-1], m[j]+intervalli[i].Peso)
					break
				}
			}
		}
	}
	fmt.Println(m)
	return m[len(intervalli)-1]
}

func maxSlice(slice []int) int {
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

func parseInput() []intervallo {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nI, _ := strconv.Atoi(sc.Text())
	intervalli := make([]intervallo, 0, nI)
	k := 0
	for k < nI {
		sc.Scan()
		var i = intervallo{0, 0, 0}
		tripla := strings.Split(sc.Text(), " ")
		i.Inizio, _ = strconv.Atoi(strings.Split(tripla[0], "-")[0])
		i.Fine, _ = strconv.Atoi(strings.Split(tripla[0], "-")[1])
		i.Peso, _ = strconv.Atoi(tripla[1])
		intervalli = append(intervalli, i)
		k++
	}
	return intervalli
}

func main() {
	intervalli := parseInput()
	sort.Slice(intervalli,
		func(i, j int) bool {
			return intervalli[i].Fine < intervalli[j].Fine
		})
	fmt.Println(sommaIntervalli(intervalli))
}
