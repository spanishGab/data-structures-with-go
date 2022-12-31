package singly

import "fmt"

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
    head *Node
    length int
}


func CreateEmptyList() SinglyLinkedList {
    return SinglyLinkedList{
        head: nil,
        length: 0,
    }
}


func Append(list *SinglyLinkedList, value int) {
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


func Push(list *SinglyLinkedList, value int) {
    var newNode Node = Node{data: value, next: nil}

    newNode.next = list.head

    list.head = &newNode

    list.length += 1
}


func Insert(list *SinglyLinkedList, value int, pos int) {
    var newNode Node = Node{data: value, next: nil}

    list.length += 1

    if list.head == nil {
        list.head = &newNode
        return
    }

    var iterator *Node = list.head

    // zero indexed pos
    for i := 0; i < pos - 1; i++ {
        iterator = iterator.next
    }

    newNode.next = iterator.next
    iterator.next = &newNode
}


func Delete(list *SinglyLinkedList) {
    if (list.head == nil) {
        return
    }

    list.length -= 1
    list.head = list.head.next
}

func Pop(list *SinglyLinkedList) {
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


func Remove(list *SinglyLinkedList, pos int) {
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


func SearchElementData(list *SinglyLinkedList, pos int) int {
    var element *Node = list.head

    for i := 0; i < pos; i++ {
        element = element.next
    }

    return element.data
}


func PrintList(list *SinglyLinkedList) {
    var iterator *Node = list.head

    fmt.Printf("list_length: %d, list_content: ", list.length)
    for iterator != nil {
        fmt.Printf("%d->", iterator.data)
        iterator = iterator.next
    }
    fmt.Println()
}
