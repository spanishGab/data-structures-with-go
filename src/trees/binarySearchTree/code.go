package binarysearchtree

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

type BinarySearchTree struct {
	root *node
}

func New() BinarySearchTree {
	return BinarySearchTree{root: nil}
}

func (tree *BinarySearchTree) IsEmpty() bool {
	return tree.root == nil
}

func isEmptyNode(root *node) bool {
	return root == nil
}

func (tree *BinarySearchTree) Height(root *node) uint {
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

func (tree *BinarySearchTree) TotalNodes(root *node) uint {
	if isEmptyNode(root) {
		return 0
	}
	var leftHeight uint = tree.TotalNodes(root.left)
	var rightHeight uint = tree.TotalNodes(root.right)
	return leftHeight + rightHeight + 1
}

func (tree *BinarySearchTree) Insert(value int) {
	var newNode node = node{
		data:  value,
		left:  nil,
		right: nil,
	}
	if tree.IsEmpty() {
		tree.root = &newNode
		return
	}

	var lastNode *node = nil
	var currentNode *node = tree.root
	for currentNode != nil {
		lastNode = currentNode
		if value == currentNode.data {
			return // there can be no repeated values
		}
		if value < currentNode.data {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}
	if value < lastNode.data {
		lastNode.left = &newNode
	} else {
		lastNode.right = &newNode
	}
}

func searchLeftTreeRightmostNode(n *node) *node {
	var lastNode *node = n
	var currentNode *node = n.left
	for currentNode.right != nil {
		lastNode = currentNode
		currentNode = currentNode.right
	}
	lastNode.right = nil
	return currentNode
}

func (tree *BinarySearchTree) Remove(value int) {
	if tree.IsEmpty() {
		return
	}

	var lastNode *node = nil
	var currentNode *node = tree.root
	for value != currentNode.data {
		lastNode = currentNode
		if value < currentNode.data {
			currentNode = currentNode.left
		} else if value > currentNode.data {
			currentNode = currentNode.right
		}
	}
	if currentNode == tree.root {
		tree.root = searchLeftTreeRightmostNode(currentNode)
		tree.root.left = currentNode.left
		tree.root.right = currentNode.right
	} else if currentNode.right != nil {
		if currentNode.left != nil {
			if lastNode.left == currentNode {
				lastNode.left = searchLeftTreeRightmostNode(currentNode)
				return
			} else {
				lastNode.right = searchLeftTreeRightmostNode(currentNode)
				return
			}
		}
		lastNode.right = currentNode.right
	} else if currentNode.left != nil {
		lastNode.left = currentNode.left
	} else {
		if lastNode.left == currentNode {
			lastNode.left = nil
		} else {
			lastNode.right = nil
		}
	}
}

func (tree *BinarySearchTree) Repr(root *node) string {
	var repr string = ""
	if isEmptyNode(root) {
		return ""
	}
	repr += tree.Repr(root.left)
	repr += fmt.Sprintf("/%d\\", root.data)
	repr += tree.Repr(root.right)
	return repr
}
