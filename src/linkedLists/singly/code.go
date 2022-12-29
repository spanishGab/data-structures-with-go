package main

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


func InsertAtEnd(list *SinglyLinkedList, value int) {
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


func InsertAtBeginning(list *SinglyLinkedList, value int) {
    var newNode Node = Node{data: value, next: nil}

    newNode.next = list.head

    list.head = &newNode

    list.length += 1
}


func InsertAtPosition(list *SinglyLinkedList, value int, pos int) {
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


func DeleteFromBegginning(list *SinglyLinkedList) {
    if (list.head == nil) {
        return
    }

    list.length -= 1
    list.head = list.head.next
}

func DeleteFromEnd(list *SinglyLinkedList) {
    if list.head == nil {
        return
    }

    var previousElement *Node = list.head
    var currentElement *Node = list.head

    list.length -= 1

    for currentElement.next != nil {
        previousElement = currentElement
        currentElement = currentElement.next
    }

    previousElement.next = nil
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


func fifoExample() {
	var list SinglyLinkedList = CreateEmptyList()

    InsertAtBeginning(&list, 10)
    PrintList(&list)

    InsertAtEnd(&list, 20)
    PrintList(&list)

    InsertAtPosition(&list, 25, 2)
    PrintList(&list)

    var iterator *Node = list.head

    for iterator != nil {
        DeleteFromBegginning(&list)
        PrintList(&list)

        iterator = iterator.next
    }
}

func lifoExample() {
	var list SinglyLinkedList = CreateEmptyList()

    InsertAtPosition(&list, 5, 0)
    PrintList(&list)

    InsertAtEnd(&list, 10)
    PrintList(&list)

    InsertAtEnd(&list, 20)
    PrintList(&list)

    var iterator *Node = list.head

    for iterator != nil {
        DeleteFromEnd(&list)
        PrintList(&list)

        iterator = iterator.next
    }
}

func main() {
    fmt.Println("Begin Fifo Example:")
    fifoExample()
    fmt.Print("End Fifo Example\n\n")

    fmt.Println("Begin Lifo Example:")
    lifoExample()
    fmt.Print("End Lifo Example\n\n")
}
