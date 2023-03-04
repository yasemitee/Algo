package main

import "fmt"

type bitreeNode struct {
	child   *bitreeNode
	sibling *bitreeNode
	val     int
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
	inorder(node.child)
	fmt.Print(node.val, " ")
	inorder(node.sibling)
}

func preorder(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, " ")
	preorder(node.child)
	preorder(node.sibling)
}
func postorder(node *bitreeNode) {
	if node == nil {
		return
	}
	postorder(node.child)
	postorder(node.sibling)
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
	if node.sibling != nil {
		stampaAlberoASommario(node.sibling, spaces)
	}
	if node.child != nil {
		stampaAlberoASommario(node.child, spaces+1)
	}
}

/*
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
*/
func arr2tree(a []int, i int) (root *bitreeNode) {
	if i >= len(a) {
		return nil
	}
	root = newNode(a[i])
	root.child = arr2tree(a, i*2+1)
	root.sibling = arr2tree(a, i+1)
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

	//printTree(albero)
	//trovaPadre(albero.root, 89, -1)
	stampaAlberoASommario(albero.root, 0)
	//stampaAlbero(t.root)
}
