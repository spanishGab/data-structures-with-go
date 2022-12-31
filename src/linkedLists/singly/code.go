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


func append(list *SinglyLinkedList, value int) {
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


func push(list *SinglyLinkedList, value int) {
    var newNode Node = Node{data: value, next: nil}

    newNode.next = list.head

    list.head = &newNode

    list.length += 1
}


func insert(list *SinglyLinkedList, value int, pos int) {
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


func delete(list *SinglyLinkedList) {
    if (list.head == nil) {
        return
    }

    list.length -= 1
    list.head = list.head.next
}

func pop(list *SinglyLinkedList) {
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


func remove(list *SinglyLinkedList, pos int) {
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

    push(&list, 5)
    PrintList(&list)

    append(&list, 10)
    PrintList(&list)

    insert(&list, 15, 2)
    PrintList(&list)

    var iterator *Node = list.head

    for iterator != nil {
        delete(&list)
        PrintList(&list)

        iterator = iterator.next
    }
}

func lifoExample() {
	var list SinglyLinkedList = CreateEmptyList()

    insert(&list, 5, 0)
    PrintList(&list)

    append(&list, 10)
    PrintList(&list)

    append(&list, 15)
    PrintList(&list)

    pop(&list)
    PrintList(&list)

    pop(&list)
    PrintList(&list)

    pop(&list)
    PrintList(&list)
}


func randomExample() {
	var list SinglyLinkedList = CreateEmptyList()

    append(&list, 10)
    PrintList(&list)

    push(&list, 5)
    PrintList(&list)

    append(&list, 15)
    PrintList(&list)

    remove(&list, 1)
    PrintList(&list)

    remove(&list, 0)
    PrintList(&list)

    remove(&list, 0)
    PrintList(&list)

}

func main() {
    fmt.Println("Begin Fifo Example:")
    fifoExample()
    fmt.Print("End Fifo Example\n\n")

    fmt.Println("Begin Lifo Example:")
    lifoExample()
    fmt.Print("End Lifo Example\n\n")

    fmt.Println("Begin Random Example:")
    randomExample()
    fmt.Print("End Random Example\n\n")
}
