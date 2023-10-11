package main

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *Node {
	current := &Node{arr[start], nil, nil}
	left := 2*start + 1
	right := 2*start + 2
	if left < size {
		current.left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		current.right = levelOrderBinaryTree(arr, right, size)
	}
	return current
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.value, " ")
	printPreOrder(n.left)
	printPreOrder(n.right)
}

func printTree(tree *Tree) {
	fmt.Print("Tree: ")
	printPreOrder(tree.root)
	fmt.Println()
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := LevelOrderBinaryTree(arr)
	printTree(tree)
}
