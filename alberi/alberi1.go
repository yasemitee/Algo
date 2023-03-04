package main

import "fmt"

type bitreeNode struct {
	left  *bitreeNode
	right *bitreeNode
	val   int
}

type bitree struct {
	root *bitreeNode
}

func newNode(val int) *bitreeNode {
	return &bitreeNode{nil, nil, val}
}
func printTree(t *bitree) {
	fmt.Print("preorder: ")
	preorder(t.root)
	fmt.Println()
	fmt.Print("inorder: ")
	inorder(t.root)
	fmt.Println()
	fmt.Print("postorder: ")
	postorder(t.root)
}

func inorder(node *bitreeNode) {
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Print(node.val, " ")
	inorder(node.right)
}

func preorder(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, " ")
	preorder(node.left)
	preorder(node.right)
}
func postorder(node *bitreeNode) {
	if node == nil {
		return
	}
	postorder(node.left)
	postorder(node.right)
	fmt.Print(node.val, " ")
}

func stampaAlberoASommario(node *bitreeNode, spaces int) {

	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	if node == nil {
		fmt.Println("*")
		return
	}
	fmt.Print("*")
	fmt.Println(node.val)

	if node.left != nil || node.right != nil {
		stampaAlberoASommario(node.left, spaces+1)
		stampaAlberoASommario(node.right, spaces+1)
	}
}

func stampaAlbero(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, "")
	if node.left == nil && node.right == nil {
		return
	}
	fmt.Print(" [")
	if node.left != nil {
		stampaAlbero(node.left)
	} else {
		fmt.Print("-")
	}
	fmt.Print(", ")
	if node.right != nil {
		stampaAlbero(node.right)
	} else {
		fmt.Print("-")
	}
	fmt.Print("]")
}

func arr2tree(a []int, i int) (root *bitreeNode) {
	if i >= len(a) {
		return nil
	}
	root = newNode(a[i])
	root.left = arr2tree(a, i*2+1)
	root.right = arr2tree(a, i*2+2)
	return root
}

func main() {
	//slice di un altro albero
	a := []int{69, 89, 28, 39, 66, 44, 12, 2, 71}
	albero := &bitree{arr2tree(a, 0)}

	printTree(albero)
}
