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

func (list *SinglyLinkedList) Append(value int) {
	var newNode Node = Node{data: value, next: nil}

	list.length += 1

	if list.head == nil {
		list.head = &newNode
		return
	}

	var iterator *Node = list.head

	for iterator.next != nil {
		iterator = iterator.next
	}

	iterator.next = &newNode
}

func (list *SinglyLinkedList) Push(value int) {
	var newNode Node = Node{data: value, next: nil}

	newNode.next = list.head

	list.head = &newNode

	list.length += 1
}

func (list *SinglyLinkedList) Insert(value int, pos int) {
	var newNode Node = Node{data: value, next: nil}

	list.length += 1

	if list.head == nil {
		list.head = &newNode
		return
	}

	var iterator *Node = list.head

	// zero indexed pos
	for i := 0; i < pos-1; i++ {
		iterator = iterator.next
	}

	newNode.next = iterator.next
	iterator.next = &newNode
}

func (list *SinglyLinkedList) Delete() {
	if list.head == nil {
		return
	}

	list.length -= 1
	list.head = list.head.next
}

func (list *SinglyLinkedList) Pop() {
	if list.head == nil {
		return
	}

	list.length -= 1

	if list.head.next == nil {
		list.head = nil
		return
	}

	var previousElement *Node
	var currentElement *Node = list.head

	for currentElement.next != nil {
		previousElement = currentElement
		currentElement = currentElement.next
	}

	previousElement.next = nil
}

func (list *SinglyLinkedList) Remove(pos int) {
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

func (list *SinglyLinkedList) Get(pos int) int {
	var element *Node = list.head

	for i := 0; i < pos; i++ {
		element = element.next
	}

	return element.data
}

func (list *SinglyLinkedList) PrintList() string {
	var iterator *Node = list.head

	var listRepr string

	listRepr = "length: " + strconv.FormatInt(int64(list.length), 10) + ", content: "
	for iterator != nil {
		listRepr += strconv.FormatInt(int64(iterator.data), 10) + "->"
		iterator = iterator.next
	}
	return listRepr
}
