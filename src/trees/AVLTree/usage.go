package avltree

import "fmt"

func Example1() {
	var tree AVLTree = New()

	tree.Insert(5)
	fmt.Println("Tree:", tree.Repr(tree.root))
  fmt.Println("Root:", tree.root.data)

	tree.Insert(10)
	fmt.Println("Tree:", tree.Repr(tree.root))
  fmt.Println("Root:", tree.root.data)

	tree.Insert(15)
	fmt.Println("Tree:", tree.Repr(tree.root))
  fmt.Println("Root:", tree.root.data)

	tree.Insert(20)
	fmt.Println("Tree:", tree.Repr(tree.root))
  fmt.Println("Root:", tree.root.data)

  tree.Insert(17)
	fmt.Println("Tree:", tree.Repr(tree.root))
  fmt.Println("Root:", tree.root.data)

	tree.Insert(-5)
	fmt.Println("Tree:", tree.Repr(tree.root))
	fmt.Println("Root:", tree.root.data)

	tree.Insert(2)
	fmt.Println("Tree:", tree.Repr(tree.root))
	fmt.Println("Root:", tree.root.data)

  tree.Insert(6)
	fmt.Println("Tree:", tree.Repr(tree.root))
	fmt.Println("Root:", tree.root.data)

  tree.Insert(21)
	fmt.Println("Tree:", tree.Repr(tree.root))
  fmt.Println("Root:", tree.root.data)

	fmt.Println("Height:", tree.Height(tree.root), "Total Nodes:", tree.TotalNodes(tree.root))

	tree.Remove(5)
	fmt.Println("Tree:", tree.Repr(tree.root), "| Tree root has changed to", tree.root.data)
	fmt.Println("Root:", tree.root.data)

	tree.Remove(-5)
	fmt.Println("Tree:", tree.Repr(tree.root))
	fmt.Println("Root:", tree.root.data)

	tree.Remove(10) // FIXME
	fmt.Println("Tree:", tree.Repr(tree.root))
	fmt.Println("Root:", tree.root.data)

  fmt.Println("Is 2 present in the tree?", tree.Has(2))
  fmt.Println("Is 5 present in the tree?", tree.Has(5))

	fmt.Println("Height:", tree.Height(tree.root), "Total Nodes:", tree.TotalNodes(tree.root))

  tree.Remove(20)
	fmt.Println("Tree:", tree.Repr(tree.root))
	fmt.Println("Root:", tree.root.data)

  tree.Remove(2)
  fmt.Println("Tree:", tree.Repr(tree.root), "| Tree root has changed to", tree.root.data)
	fmt.Println("Root:", tree.root.data)

  tree.Remove(15)
	fmt.Println("Tree:", tree.Repr(tree.root))
	fmt.Println("Root:", tree.root.data)

  tree.Remove(6)
	fmt.Println("Tree:", tree.Repr(tree.root))

  tree.Remove(17)
	fmt.Println("Tree:", tree.Repr(tree.root))

  tree.Remove(21)
	fmt.Println("Tree:", tree.Repr(tree.root))

  fmt.Println("Height:", tree.Height(tree.root), "Total Nodes:", tree.TotalNodes(tree.root))
}
