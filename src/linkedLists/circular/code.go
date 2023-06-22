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

func (list *CircularLinkedList) Delete(pos int) {
  if list.head == nil && list.tail == nil {
		return
	}

  list.length -= 1

  if pos == 0 {
    if list.length == 0 {
      list.head = nil
      list.tail = nil
      return
    }
    list.head = list.head.next
    list.head.prev = list.tail
    list.tail.next = list.head
    return
  } else if pos == list.length {
    list.tail = list.tail.prev
    list.tail.next = list.head
    list.head.prev = list.tail
    return
  }

  var previousElement *Node
  var currentElement *Node = list.head

  for i := 0; i < pos; i++ {
    previousElement = currentElement
    currentElement = currentElement.next
  }

  previousElement.next = currentElement.next
  currentElement.next.prev = previousElement
}

func (list *CircularLinkedList) Pop() {
  list.Delete(list.length - 1)
}

func (list *CircularLinkedList) Get(pos int) int {
  var element *Node = list.head

  if pos > list.length - 1 {
    panic("List index out of range")
  } else if pos == list.length - 1 {
    return list.tail.data
  }

  for i := 0; i < pos; i++ {
    element = element.next
  }
  return element.data
}

func (list *CircularLinkedList) Sort() {
  if list.head == nil || list.head.next == nil {
    return
  }

  var key int
  var previousElement *Node
  var index *Node = list.head.next
  for index != list.head {
    key = index.data
    previousElement = index.prev
    for previousElement != list.tail {
      if key > previousElement.data {
        break
      }
      previousElement.next.data = previousElement.data
      previousElement.data = key
      previousElement = previousElement.prev
    }
    index = index.next
  }
}

func (list *CircularLinkedList) Reverse() {
  if list.head == nil {
    return
  }

  var index *Node = list.head
  var aux *Node
  for index.next != list.head {
    aux = index.prev
    index.prev = index.next
    index.next = aux
    index = index.prev
  }
  list.tail = index.next
  aux = index.prev
  index.prev = list.tail
  index.next = aux
  list.head = index
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
