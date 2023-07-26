package binarytree

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func New() BinaryTree {
	return BinaryTree{root: nil}
}

func IsEmpty(root *Node) bool {
  return root == nil
}

func Height(root *Node) uint {
  if IsEmpty(root) {
    return 0
  }
  var leftHeight uint = Height(root.left)
  var rightHeight uint = Height(root.right)

  if leftHeight > rightHeight {
    return leftHeight + 1
  }
  return rightHeight + 1
}

func TotalNodes(root *Node) uint {
  if IsEmpty(root) {
    return 0
  }
  var leftHeight uint = TotalNodes(root.left)
  var rightHeight uint = TotalNodes(root.right)
  return leftHeight + rightHeight + 1
}

func Repr(root *Node) string {
  var repr string = "Post-Ordered Tree: "
  if IsEmpty(root) {
    return ""
  }
  Repr(root.left)
  Repr(root.right)
  repr += fmt.Sprintf("%d", root.data)
  return repr
}
