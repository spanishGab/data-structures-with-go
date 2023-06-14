package circular

import (
	"strconv"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type CircularLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func New() CircularLinkedList {
	return CircularLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (list *CircularLinkedList) Insert(value int, pos int) {
	var newNode Node = Node{data: value, next: nil, prev: nil}

	list.length += 1

	if list.head == nil && list.tail == nil {
		list.head = &newNode
		list.tail = &newNode
		list.head.next = list.tail
		list.head.prev = list.tail
		list.tail.next = list.head
		list.tail.prev = list.head
		return
	} else if pos == 0 {
		newNode.next = list.head
		newNode.prev = list.tail
		list.head.prev = &newNode
		list.tail.next = &newNode
		list.head = &newNode
		return
	} else if pos == list.length-1 {
		newNode.prev = list.tail
		newNode.next = list.head
		list.tail.next = &newNode
		list.head.prev = &newNode
		list.tail = &newNode
		return
	}

	var iterator *Node = list.head

	// zero indexed pos
	for i := 1; i < pos; i++ {
		iterator = iterator.next
	}
	newNode.next = iterator.next
	newNode.prev = iterator
	newNode.next.prev = &newNode
	iterator.next = &newNode
}

func (list *CircularLinkedList) Append(value int) {
	list.Insert(value, list.length)
}

func (list *CircularLinkedList) Repr() string {
	var iterator *Node = list.head

	var listRepr string = ""

	listRepr = "length: " + strconv.FormatInt(int64(list.length), 10) + ", content: tail<->"
	for i := 0; i < list.length; i++ {
		listRepr += strconv.FormatInt(int64(iterator.data), 10) + "<->"
		iterator = iterator.next
	}
	listRepr += "head"
	return listRepr
}
