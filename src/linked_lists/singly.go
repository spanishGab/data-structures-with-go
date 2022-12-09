package main

type Node struct {
	data int
	next *Node
}

func createNode(value int, nextNode Node) Node {
	return Node{data: value, next: &nextNode}
}

func main() {
	var newNode Node
	newNode.data = 1
	newNode.next = nil
}
