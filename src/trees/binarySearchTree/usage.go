package binarysearchtree

import "fmt"

func Example1() {
	var tree BinarySearchTree = New()

	tree.Insert(5)
	fmt.Println("Tree:", tree.Repr(tree.root))

	tree.Insert(10)
	fmt.Println("Tree:", tree.Repr(tree.root))

	tree.Insert(0)
	fmt.Println("Tree:", tree.Repr(tree.root))

	tree.Insert(15)
	fmt.Println("Tree:", tree.Repr(tree.root))

	tree.Insert(-5)
	fmt.Println("Tree:", tree.Repr(tree.root))

	tree.Insert(2)
	fmt.Println("Tree:", tree.Repr(tree.root))
	fmt.Println("Height:", tree.Height(tree.root), "Total Nodes:", tree.TotalNodes(tree.root))

	tree.Remove(5)
	fmt.Println("Tree:", tree.Repr(tree.root))

	tree.Remove(-5)
	fmt.Println("Tree:", tree.Repr(tree.root))

	tree.Remove(10)
	fmt.Println("Tree:", tree.Repr(tree.root))
	fmt.Println("Height:", tree.Height(tree.root), "Total Nodes:", tree.TotalNodes(tree.root))
}
