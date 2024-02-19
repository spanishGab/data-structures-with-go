package avltree

import "fmt"

func Example1() {
	var tree AVLTree = New()

	tree.Insert(5)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
  fmt.Printf("Root: %d\n\n", tree.root.data)

	tree.Insert(10)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
  fmt.Printf("Root: %d\n\n", tree.root.data)

	tree.Insert(15)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
  fmt.Printf("Root: %d\n\n", tree.root.data)

	tree.Insert(20)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
  fmt.Printf("Root: %d\n\n", tree.root.data)

  tree.Insert(17)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
  fmt.Printf("Root: %d\n\n", tree.root.data)

	tree.Insert(-5)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
	fmt.Printf("Root: %d\n\n", tree.root.data)

	tree.Insert(2)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
	fmt.Printf("Root: %d\n\n", tree.root.data)

  tree.Insert(6)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
	fmt.Printf("Root: %d\n\n", tree.root.data)

  tree.Insert(11)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
  fmt.Printf("Root: %d\n\n", tree.root.data)

	fmt.Println("Height:", tree.Height(tree.root))

	tree.Remove(5)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
	fmt.Printf("Root: %d\n\n", tree.root.data)

	tree.Remove(-5)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
	fmt.Printf("Root: %d\n\n", tree.root.data)

	tree.Remove(10)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
	fmt.Printf("Root: %d\n\n", tree.root.data)

	fmt.Println("Height:", tree.Height(tree.root))

  tree.Remove(20)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
	fmt.Printf("Root: %d\n\n", tree.root.data)

  tree.Remove(2)
  fmt.Println("Tree:\n", tree.Repr(tree.root))
	fmt.Printf("Root: %d\n\n", tree.root.data)

  tree.Remove(15)
	fmt.Println("Tree:\n", tree.Repr(tree.root))
	fmt.Printf("Root: %d\n\n", tree.root.data)

  tree.Remove(6)
	fmt.Println("Tree:\n", tree.Repr(tree.root))

  tree.Remove(17)
	fmt.Println("Tree:\n", tree.Repr(tree.root))

  tree.Remove(11)
	fmt.Println("Tree:\n", tree.Repr(tree.root))

  fmt.Println("Height:", tree.Height(tree.root))
}
