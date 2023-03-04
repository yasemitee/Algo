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
	/* t := &bitree{nil}
	t.root = &bitreeNode{nil, nil, 78}
	//inizio sotto alberso sx
	t.root.left = newNode(54)
	t.root.left.right = newNode(90)
	t.root.left.right.left = newNode(19)
	t.root.left.right.right = newNode(95)
	//inizio sotto albero dx
	t.root.right = newNode(21)
	t.root.right.left = newNode(16)
	t.root.right.right = newNode(19)
	t.root.right.left.left = newNode(5)
	t.root.right.right.left = newNode(56)
	t.root.right.right.right = newNode(43) */

	/* t2 := &bitree{nil}
	t2.root = &bitreeNode{nil, nil, 10}
	t2.root.left = newNode(5)
	t2.root.left.left = newNode(3)
	t2.root.left.right = newNode(2)
	t2.root.right = newNode(6)
	t2.root.right.left = newNode(21) */

	//slice di un altro albero
	a := []int{69, 89, 28, 39, 66, 44, 12, 2, 71}
	albero := &bitree{arr2tree(a, 0)}

	printTree(albero)
	//trovaPadre(albero.root, 89, -1)
	//stampaAlberoASommario(t.root, 0)
	//stampaAlbero(t.root)
}
