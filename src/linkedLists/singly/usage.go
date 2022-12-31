package singly

import "fmt"


func Example1() {
	var list SinglyLinkedList = CreateEmptyList()

    Push(&list, 5)
    PrintList(&list)

    Append(&list, 10)
    PrintList(&list)

    Insert(&list, 15, 2)
    PrintList(&list)

    var iterator *Node = list.head

    for iterator != nil {
        Delete(&list)
        PrintList(&list)

        iterator = iterator.next
    }
}

func Example2() {
	var list SinglyLinkedList = CreateEmptyList()

    Insert(&list, 5, 0)
    PrintList(&list)

    Append(&list, 10)
    PrintList(&list)

    Append(&list, 15)
    PrintList(&list)

    Pop(&list)
    PrintList(&list)

    Pop(&list)
    PrintList(&list)

    Pop(&list)
    PrintList(&list)
}


func Example3() {
	var list SinglyLinkedList = CreateEmptyList()

    Append(&list, 10)
    PrintList(&list)

    Push(&list, 5)
    PrintList(&list)

    Append(&list, 15)
    PrintList(&list)

    Remove(&list, 1)
    PrintList(&list)

    Remove(&list, 0)
    PrintList(&list)

    Remove(&list, 0)
    PrintList(&list)

}


func Example4() {
	var list SinglyLinkedList = CreateEmptyList()

    Append(&list, 10)
    PrintList(&list)

    Push(&list, 5)
    PrintList(&list)

    Append(&list, 15)
    PrintList(&list)

    for i := 0; i < list.length; i++ {
        fmt.Printf("Element at position %d: %d\n", i, SearchElementData(&list, i))
    }
}
