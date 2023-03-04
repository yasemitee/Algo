package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
}

func newNode(val int) *listNode {
	return &listNode{val, nil}
}

func addNewNode(l *linkedList, val int) bool {
	flag, _ := searchList(*l, val)
	if flag {
		return false
	}
	node := newNode(val)
	node.next = l.head
	l.head = node
	return true
}

func addNewNodeSorted(l *linkedList, val int) bool {
	var curr, prev *listNode = l.head, nil
	if l.head == nil {
		l.head = newNode(val)
		return true
	}
	for curr != nil {
		if curr.item > val {
			if prev == nil {
				node := newNode(val)
				node.next = l.head
				l.head = node
				return true
			} else {
				node := newNode(val)
				prev.next = node
				node.next = curr
				return true
			}
		} else if curr.next == nil && curr.item < val {
			node := newNode(val)
			curr.next = node
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}

func printList(l linkedList) {
	p := l.head
	for p != nil {
		fmt.Print(p.item, " ")
		p = p.next
	}
	fmt.Println()
}

func reversePrintList(l linkedList) {
	p := l.head
	aList := make([]int, 0)
	for p != nil {
		aList = append(aList, p.item)
		p = p.next
	}
	for k := len(aList) - 1; k >= 0; k-- {
		fmt.Print(aList[k], " ")
	}
	fmt.Println()
}

func searchList(l linkedList, val int) (bool, *listNode) {
	p := l.head
	for p != nil {
		if p.item == val {
			return true, p
		}
		p = p.next
	}
	return false, nil
}

func removeItem(l *linkedList, val int) bool {
	var curr, prev *listNode = l.head, nil
	for curr != nil {
		if curr.item == val {
			if prev == nil {
				l.head = l.head.next
			} else {
				prev.next = curr.next
			}
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}

func removeAll(l *linkedList) {
	l.head = nil
}

func listLength(l linkedList) int {
	var count int
	p := l.head
	for p != nil {
		count++
		p = p.next
	}
	return count
}

func main() {
	var input string
	var l linkedList
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
		if input == "f" {
			return
		}
		//divido l'operazione dal valore
		var op string = string(input[0])
		var val int
		val, _ = strconv.Atoi(string(input[1:]))
		//operszione in base alla scelta data in input
		switch op {
		case "+":
			flag := addNewNode(&l, val)
			if flag == false {
				fmt.Println("elemento già presente nella lista")
			}
		case "s":
			flag := addNewNodeSorted(&l, val)
			if flag == true {
				fmt.Println("elemento inserito nella posizione giusta")
			}
		case "-":
			flag := removeItem(&l, val)
			if flag {
				fmt.Println("elemento eliminato con successo")
			} else {
				fmt.Println("elemento non presente nella lista")
			}
		case "?":
			flag, pos := searchList(l, val)
			if flag {
				fmt.Println("il numero:", val, "si trova nella lista in posizione", pos)
			} else {
				fmt.Println("numero non trovato")
			}
		case "c":
			fmt.Println("numero di elementi della list:", listLength(l))
		case "p":
			printList(l)
		case "o":
			reversePrintList(l)
		case "d":
			removeAll(&l)
		default:
			fmt.Println("comando non valido")
		}
	}
}
