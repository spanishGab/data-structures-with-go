package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
    head *Node
    iterator *Node
    length int
}


func CreateEmptyList() SinglyLinkedList {
    return SinglyLinkedList{
        head: nil,
        iterator: nil,
        length: 0,
    }
}


func InsertAtEnd(list *SinglyLinkedList, value int) {
    var newNode Node = Node{data: value, next: nil}

    if list.head == nil {
        list.head = &newNode
    } else {
        resetIterator(list)

        for list.iterator.next != nil {
            list.iterator = list.iterator.next
        }

        list.iterator.next = &newNode
    }

    resetIterator(list)
}


func InsertAtBeginning(list *SinglyLinkedList, value int) {
    var newNode Node = Node{data: value, next: nil}

    if list.head != nil {
        newNode.next = list.head
    }

    list.head = &newNode

    resetIterator(list)
}


func InsertAtPosition(list *SinglyLinkedList, value int, pos int) {
    var newNode Node = Node{data: value, next: nil}

    if list.head == nil {
        list.head = &newNode
    } else {
        resetIterator(list)

        // zero indexed pos
        for i := 0; i < pos - 1; i++ {
            list.iterator = list.iterator.next
        }

        newNode.next = list.iterator.next
        list.iterator.next = &newNode
    }

    resetIterator(list)
}


func resetIterator(list *SinglyLinkedList) {
    list.iterator = list.head
}

func PrintList(list *SinglyLinkedList) {
    for list.iterator != nil {
        fmt.Printf("%d ", list.iterator.data)
        list.iterator = list.iterator.next
    }
    fmt.Println()

    resetIterator(list)
}


func main() {
	var list SinglyLinkedList = CreateEmptyList()

    InsertAtEnd(&list, 15)
    PrintList(&list)


    InsertAtBeginning(&list, 10)
    PrintList(&list)

    InsertAtEnd(&list, 20)
    PrintList(&list)

    InsertAtBeginning(&list, 5)
    PrintList(&list)

    InsertAtEnd(&list, 30)
    PrintList(&list)

    InsertAtPosition(&list, 25, 4)
    PrintList(&list)
}
