package doubly

import (
	"fmt"
)

func Example1() {
	var list DoublyLinkedList = New()

	list.Insert(5, 0)
	fmt.Println(list.Repr())

	list.Append(10)
	fmt.Println(list.Repr())

	list.Insert(15, 0)
	fmt.Println(list.Repr())

  fmt.Println("Get:", list.Get(0), list.Get(1), list.Get(2))

  list.Sort()
	fmt.Println(list.Repr())

  list.Reverse()
	fmt.Println(list.Repr())

  for list.head != nil {
		list.Delete(0)
		fmt.Println(list.Repr())
	}
}

func Example2() {
	var list DoublyLinkedList = New()

	list.Insert(5, 0)
	fmt.Println(list.Repr())

	list.Sort()
	fmt.Println(list.Repr())

	list.Append(10)
	fmt.Println(list.Repr())

	list.Append(15)
	fmt.Println(list.Repr())

  list.Reverse()
	fmt.Println(list.Repr())

	for list.head != nil {
		list.Pop()
		fmt.Println(list.Repr())
	}
}

func Example3() {
	var list DoublyLinkedList = New()

	list.Append(10)
	fmt.Println(list.Repr())

	list.Insert(5, 0)
	fmt.Println(list.Repr())

	list.Reverse()
	fmt.Println(list.Repr())

	list.Append(15)
	fmt.Println(list.Repr())

	list.Delete(1)
	fmt.Println(list.Repr())

	list.Pop()
	fmt.Println(list.Repr())

	list.Delete(0)
	fmt.Println(list.Repr())
}

func Example4() {
	var list DoublyLinkedList = New()

	list.Append(10)
	fmt.Println(list.Repr())

	list.Insert(5, 1)
	fmt.Println(list.Repr())

	list.Sort()
	fmt.Println(list.Repr())
}

func Example5() {
	var list DoublyLinkedList = New()

	list.Append(10)
	fmt.Println(list.Repr())

	list.Append(5)
	fmt.Println(list.Repr())

	list.Append(15)
	fmt.Println(list.Repr())

	list.Append(20)
	fmt.Println(list.Repr())

	list.Append(0)
	fmt.Println(list.Repr())

	list.Reverse()
	fmt.Println(list.Repr())
}
