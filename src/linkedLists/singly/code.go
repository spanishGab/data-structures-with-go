package singly

import (
	"strconv"
)

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head   *Node
	length int
}

func New() SinglyLinkedList {
	return SinglyLinkedList{
		head:   nil,
		length: 0,
	}
}

func (list *SinglyLinkedList) Insert(value int, pos int) {
	var newNode Node = Node{data: value, next: nil}

	list.length += 1

	if list.head == nil {
		list.head = &newNode
		return
	} else if pos == 0 {
		newNode.next = list.head
		list.head = &newNode
		return
	}

	var iterator *Node = list.head

	// zero indexed pos
	for i := 1; i < pos; i++ {
		iterator = iterator.next
	}

	newNode.next = iterator.next
	iterator.next = &newNode
}

func (list *SinglyLinkedList) Append(value int) {
	list.Insert(value, list.length)
}

func (list *SinglyLinkedList) Delete(pos int) {
	if list.head == nil {
		return
	}

	list.length -= 1

	if list.head.next == nil {
		list.head = nil
		return
	}

	if pos == 0 {
		list.head = list.head.next
		return
	}

	var previousElement *Node
	var currentElement *Node = list.head

	for i := 0; i < pos; i++ {
		previousElement = currentElement
		currentElement = currentElement.next
	}

	previousElement.next = currentElement.next
}

func (list *SinglyLinkedList) Pop() {
	list.Delete(list.length - 1)
}

func (list *SinglyLinkedList) Get(pos int) int {
	var element *Node = list.head

	for i := 0; i < pos; i++ {
		element = element.next
	}

	return element.data
}

func (list *SinglyLinkedList) Sort() {
	if list.head.next == nil {
		return
	}

	var currentElement *Node = list.head
	var index *Node
	var aux int
	for currentElement != nil {
		index = currentElement.next
		for index != nil {
			if currentElement.data > index.data {
				aux = index.data
				index.data = currentElement.data
				currentElement.data = aux
			}
			index = index.next
		}
		currentElement = currentElement.next
	}
}

func (list *SinglyLinkedList) Repr() string {
	var iterator *Node = list.head

	var listRepr string = ""

	listRepr = "length: " + strconv.FormatInt(int64(list.length), 10) + ", content: "
	for iterator != nil {
		listRepr += strconv.FormatInt(int64(iterator.data), 10) + "->"
		iterator = iterator.next
	}
	return listRepr
}
