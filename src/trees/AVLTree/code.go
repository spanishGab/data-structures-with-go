package avltree

import (
	"fmt"
)

type node struct {
	data   int
  height int
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

func (n *node) isEmptyNode() bool {
	return n == nil
}

func (tree *AVLTree) Height(root *node) uint {
	if root.isEmptyNode() {
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
	if root.isEmptyNode() {
		return 0
	}
	var leftHeight uint = tree.TotalNodes(root.left)
	var rightHeight uint = tree.TotalNodes(root.right)
	return leftHeight + rightHeight + 1
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
	if root.isEmptyNode() {
		return ""
	}
	repr += tree.Repr(root.left)
	repr += fmt.Sprintf("/%d\\", root.data)
	repr += tree.Repr(root.right)
	return repr
}

func (tree *AVLTree) Insert(value int) {
  if tree.IsEmpty() {
    tree.root = &node{data: value}
    return
  }
  tree.root = tree.root.insertNode(value)
}

func (tree *AVLTree) Remove(value int) {
  if tree.IsEmpty() {
    return
  }
  tree.root = tree.root.removeNode(value)
}

// This methdod differs from BinaryTree as it balances the tree
// and then inserts a new item on it
func (currentNode *node) insertNode(value int) *node {
  if currentNode == nil {
    return &node{data: value}
  }
  if currentNode.data == value {
    panic("A binary tree can not have repeated values.")
  }

  if value < currentNode.data {
    currentNode.left = currentNode.left.insertNode(value)
    if currentNode.calculateBalanceFactor() >= 2 {
      if value < currentNode.left.data {
        currentNode = currentNode.rotateRight()
      } else {
        currentNode = currentNode.rotateRightLeft()
      }
    }
  } else {
    currentNode.right = currentNode.right.insertNode(value)
    if currentNode.calculateBalanceFactor() >= 2 {
      if value > currentNode.right.data {
        currentNode = currentNode.rotateLeft()
      } else {
        currentNode = currentNode.rotateLeftRight()
      }
    }
  }
  currentNode.height = biggest(getNodeHeight(currentNode.left), getNodeHeight(currentNode.right)) + 1
  return currentNode
}

// This methdod differs from BinaryTree as it balances the tree
// and then removes an item from it
func (currentNode *node) removeNode(value int) *node {
  if currentNode == nil {
    return nil
  }

  if value < currentNode.data {
    currentNode.left = currentNode.left.removeNode(value)
    if currentNode.calculateBalanceFactor() >= 2 {
      if getNodeHeight(currentNode.right.left) <= getNodeHeight(currentNode.right.right) {
        currentNode = currentNode.rotateLeft()
      } else {
        currentNode = currentNode.rotateLeftRight()
      }
    }
  } else if value > currentNode.data {
    currentNode.right = currentNode.right.removeNode(value)
    if currentNode.calculateBalanceFactor() >= 2 {
      if getNodeHeight(currentNode.left.left) <= getNodeHeight(currentNode.left.right) {
        currentNode = currentNode.rotateRight()
      } else {
        currentNode = currentNode.rotateRightLeft()
      }
    }
  } else if value == currentNode.data {
    if currentNode.left == nil {
      currentNode = currentNode.right
    } else if currentNode.right == nil {
      currentNode = currentNode.left
    } else {
      newNodeStructure, lowestValue := currentNode.right.findAndDeleteLowestElement()
      currentNode.right = newNodeStructure
      currentNode.data = lowestValue
    }
  }
  if currentNode != nil {
    currentNode.height = biggest(getNodeHeight(currentNode.left), getNodeHeight(currentNode.right)) + 1
  }
  return currentNode
}

func (n *node) findAndDeleteLowestElement() (*node, int){
	if n.left.isEmptyNode() {
    var lowestValue = n.data
    if n.right.isEmptyNode() {
      return nil, lowestValue
    }
    return n.right, lowestValue
  }

  var lastNode *node = n
	var currentNode *node = n.left
	for currentNode.left != nil {
    lastNode = currentNode
		currentNode = currentNode.left
	}
  lastNode.left = nil
	return n, currentNode.data
}

func (n *node) calculateBalanceFactor() int {
  return absSubtratcion(getNodeHeight(n.left), getNodeHeight(n.right))
}

func (rootNode *node) rotateRight() *node {
  var newNode node = node{}
  newNode = *rootNode.left
  rootNode.left = newNode.right
  newNode.right = rootNode
  rootNode.height = biggest(getNodeHeight(rootNode.left), getNodeHeight((rootNode.right))) + 1
  newNode.height = biggest((getNodeHeight(newNode.left)), rootNode.height) + 1
  return &newNode
}

func (rootNode *node) rotateLeft() *node {
  var newNode node = node{}
  newNode = *rootNode.right
  rootNode.right = newNode.left
  newNode.left = rootNode
  rootNode.height = biggest(getNodeHeight(rootNode.right), getNodeHeight((rootNode.left))) + 1
  newNode.height = biggest((getNodeHeight(newNode.right)), rootNode.height) + 1
  return &newNode
}

func (rootNode *node) rotateRightLeft() *node {
  rootNode.left = rootNode.left.rotateLeft()
  return rootNode.rotateRight()
}

func (rootNode *node) rotateLeftRight() *node {
  rootNode.right = rootNode.right.rotateRight()
  return rootNode.rotateLeft()
}

func getNodeHeight(n *node) int {
  if n == nil {
    return -1
  }
  return n.height
}

func biggest(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func absSubtratcion(a, b int) int {
  if a > b {
    return a - b
  }
  return b - a
}
