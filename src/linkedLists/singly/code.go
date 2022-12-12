package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
    head Node
    iterator *Node
    length int
}


func createEmptyList() SinglyLinkedList {
    return SinglyLinkedList{
        head: Node{data: 0, next: nil},
        iterator: &Node{data: 0, next: nil},
        length: 0,
    }
}


func insertAtEnd(list *SinglyLinkedList, value int) {
    var newNode Node = Node{data: value, next: nil}

    fmt.Println("insertAtEnd")

    if list.head.next == nil {
        fmt.Println("Criando segundo item")
        list.head.next = &newNode
        fmt.Println(list.head, list.head.next)
    } else {
        list.iterator = list.head.next

        for list.iterator.next != nil {
            list.iterator = list.iterator.next
        }

        list.iterator.next = &newNode
    }

    list.iterator = &list.head
}

func main() {
	var list SinglyLinkedList = createEmptyList()

    fmt.Println(list.head, list.iterator)
    insertAtEnd(&list, 15)


    for list.iterator != nil {
        fmt.Println(list.iterator.data)
        list.iterator = list.iterator.next
    }
}
