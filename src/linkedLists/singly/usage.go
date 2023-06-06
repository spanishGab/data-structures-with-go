package singly

import "fmt"

func Example1() {
	var list SinglyLinkedList = New()

	list.Push(5)
	fmt.Println(list.PrintList())

	list.Append(10)
	fmt.Println(list.PrintList())

	list.Insert(15, 2)
	fmt.Println(list.PrintList())

	var iterator *Node = list.head

	for iterator != nil {
		list.Delete()
		fmt.Println(list.PrintList())

		iterator = iterator.next
	}
}

func Example2() {
	var list SinglyLinkedList = New()

	list.Insert(5, 0)
	fmt.Println(list.PrintList())

	list.Append(10)
	fmt.Println(list.PrintList())

	list.Append(15)
	fmt.Println(list.PrintList())

	list.Pop()
	fmt.Println(list.PrintList())

	list.Pop()
	fmt.Println(list.PrintList())

	list.Pop()
	fmt.Println(list.PrintList())
}

func Example3() {
	var list SinglyLinkedList = New()

	list.Append(10)
	fmt.Println(list.PrintList())

	list.Push(5)
	fmt.Println(list.PrintList())

	list.Append(15)
	fmt.Println(list.PrintList())

	list.Remove(1)
	fmt.Println(list.PrintList())

	list.Remove(0)
	fmt.Println(list.PrintList())

	list.Remove(0)
	fmt.Println(list.PrintList())

}

func Example4() {
	var list SinglyLinkedList = New()

	list.Append(10)
	fmt.Println(list.PrintList())

	list.Push(5)
	fmt.Println(list.PrintList())

	list.Append(15)
	fmt.Println(list.PrintList())
}
