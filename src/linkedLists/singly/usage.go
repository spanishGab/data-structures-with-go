package singly

import "fmt"

func Example1() {
	var list SinglyLinkedList = New()

	list.Insert(5, 0)
	fmt.Println(list.Repr())

	list.Append(10)
	fmt.Println(list.Repr())

	list.Insert(15, 0)
	fmt.Println(list.Repr())

	for list.head != nil {
		list.Delete(0)
		fmt.Println(list.Repr())
	}
}

func Example2() {
	var list SinglyLinkedList = New()

	list.Insert(5, 0)
	fmt.Println(list.Repr())

	list.Append(10)
	fmt.Println(list.Repr())

	list.Append(15)
	fmt.Println(list.Repr())

	for list.head != nil {
		list.Pop()
		fmt.Println(list.Repr())
	}
}

func Example3() {
	var list SinglyLinkedList = New()

	list.Append(10)
	fmt.Println(list.Repr())

	list.Insert(5, 0)
	fmt.Println(list.Repr())

	list.Append(15)
	fmt.Println(list.Repr())

	list.Delete(1)
	fmt.Println(list.Repr())

	list.Delete(0)
	fmt.Println(list.Repr())

	list.Delete(0)
	fmt.Println(list.Repr())

}

func Example4() {
	var list SinglyLinkedList = New()

	list.Append(10)
	fmt.Println(list.Repr())

	list.Insert(5, 0)
	fmt.Println(list.Repr())

	list.Append(15)
	fmt.Println(list.Repr())
}
