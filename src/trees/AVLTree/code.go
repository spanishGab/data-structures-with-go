package avltree

import (
	"fmt"
)

type node struct {
	data   int
  height uint
	left   *node
	right  *node
}

type AVLTree struct {
	root *node
}

func New() AVLTree {
	return AVLTree{root: nil}
}

func (tree *AVLTree) IsEmpty() bool {
	return tree.root == nil
}

func isEmptyNode(root *node) bool {
	return root == nil
}

func (tree *AVLTree) Height(root *node) uint {
	if isEmptyNode(root) {
		return 0
	}
	var leftHeight uint = tree.Height(root.left)
	var rightHeight uint = tree.Height(root.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (tree *AVLTree) TotalNodes(root *node) uint {
	if isEmptyNode(root) {
		return 0
	}
	var leftHeight uint = tree.TotalNodes(root.left)
	var rightHeight uint = tree.TotalNodes(root.right)
	return leftHeight + rightHeight + 1
}


func findReplaceableElement(n *node) *node {
	var lastNode *node = n
	var currentNode *node = n.left
	for currentNode.right != nil {
    lastNode = currentNode
		currentNode = currentNode.right
	}
	lastNode.right = nil
	return currentNode
}

func (tree *AVLTree) Has(value int) bool {
  if tree.IsEmpty() {
    return false
  }

  var currentNode *node = tree.root
  for currentNode != nil {
    if currentNode.data == value {
      return true
      } else if currentNode.data < value {
        currentNode = currentNode.left
        } else {
          currentNode = currentNode.right
        }
  }
  return false
}

func (tree *AVLTree) Repr(root *node) string {
  var repr string = ""
  if tree.IsEmpty() {
    return "Tree is empty!"
  }
	if isEmptyNode(root) {
		return ""
	}
	repr += tree.Repr(root.left)
	repr += fmt.Sprintf("/%d\\", root.data)
	repr += tree.Repr(root.right)
	return repr
}

// This methdod differs from BinaryTree as it balances the tree
// and then inserts a new item on it
func (tree *AVLTree) Insert(value int) {}

// This methdod differs from BinaryTree as it balances the tree
// and then removes an item from it
func (tree *AVLTree) Remove(value int) {}

func balanceFactor(n *node) uint {
  return n.left.height - n.right.height
}
