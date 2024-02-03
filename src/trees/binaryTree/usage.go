package binarytree

import "fmt"

func Example1() {
	var tree BinaryTree = New()

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
	fmt.Println("Tree:", tree.Repr(tree.root), "| Tree root has changed to", tree.root.data)

	tree.Remove(-5)
	fmt.Println("Tree:", tree.Repr(tree.root))

	tree.Remove(10)
	fmt.Println("Tree:", tree.Repr(tree.root))

  fmt.Println("Is 2 present in the tree?", tree.Has(2))
  fmt.Println("Is 5 present in the tree?", tree.Has(5))

	fmt.Println("Height:", tree.Height(tree.root), "Total Nodes:", tree.TotalNodes(tree.root))

  tree.Remove(0)
	fmt.Println("Tree:", tree.Repr(tree.root))
  tree.Remove(2)
  fmt.Println("Tree:", tree.Repr(tree.root), "| Tree root has changed to", tree.root.data)
  tree.Remove(15)
	fmt.Println("Tree:", tree.Repr(tree.root))

  fmt.Println("Height:", tree.Height(tree.root), "Total Nodes:", tree.TotalNodes(tree.root))
}
