package main

type Node struct {
	data int
	next *Node
}

func main() {
	var newNode Node
	newNode.data = 1
	newNode.next = nil
}
