package doubly

import (
	"strconv"
)

type Node struct {
  data int
  next *Node
  prev *Node
}

type DoublyLinkedList struct {
  head *Node
  length int
}

func New() DoublyLinkedList {
  return DoublyLinkedList{
    head: nil,
    length: 0,
  }
}

func (list *DoublyLinkedList) Insert(value int, pos int) {
  var newNode Node = Node{data: value, next: nil, prev: nil}

  list.length += 1

  if list.head == nil {
    list.head = &newNode
    return
  } else if pos == 0 {
    newNode.next = list.head
    list.head.prev = &newNode
    list.head = &newNode
    return
  }

  var iterator *Node = list.head

  // zero indexed pos
  for i := 1; i < pos; i++ {
    iterator = iterator.next
  }
  newNode.next = iterator.next
  newNode.prev = iterator
  if iterator.next != nil {
    newNode.next.prev = &newNode
  }
  iterator.next = &newNode
}

func (list *DoublyLinkedList) Append(value int) {
  list.Insert(value, list.length)
}

func (list *DoublyLinkedList) Delete(pos int) {
  if list.head == nil {
		return
	}

  list.length -= 1

  if list.head.next == nil {
    list.head = nil
    return
  } else if pos == 0 {
    list.head = list.head.next
    list.head.prev = nil
    return
  }

  var previousElement *Node
  var currentElement *Node = list.head

  for i := 0; i < pos; i++ {
    previousElement = currentElement
    currentElement = currentElement.next
  }

  previousElement.next = currentElement.next
  if currentElement.next != nil {
    currentElement.next.prev = previousElement
  }
}

func (list *DoublyLinkedList) Pop() {
  list.Delete(list.length-1)
}

func (list *DoublyLinkedList) Get(pos int) int {
  var element *Node = list.head

  for i := 0; i < pos; i++ {
    element = element.next
  }
  return element.data
}

func (list *DoublyLinkedList) Sort() {
  if list.head == nil || list.head.next == nil {
    return
  }

  var key int
  var previousElement *Node
  var index *Node = list.head.next
  for index != nil {
    key = index.data
    previousElement = index.prev
    for previousElement != nil && key < previousElement.data {
      previousElement.next.data = previousElement.data
      previousElement.data = key
      previousElement = previousElement.prev
    }
    index = index.next
  }
}

func (list *DoublyLinkedList) Reverse() {
  if list.head == nil {
    return
  }

  var index *Node = list.head
  var aux *Node
  for index.next != nil {
    aux = index.prev
    index.prev = index.next
    index.next = aux
    index = index.prev
  }
  aux = index.prev
  index.prev = index.next
  index.next = aux
  list.head = index
}

func (list *DoublyLinkedList) Repr() string {
	var iterator *Node = list.head

	var listRepr string = ""

	listRepr = "length: " + strconv.FormatInt(int64(list.length), 10) + ", content: "
	for iterator != nil {
		listRepr += strconv.FormatInt(int64(iterator.data), 10) + "<->"
		iterator = iterator.next
	}
	return listRepr
}
