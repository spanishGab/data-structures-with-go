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

  for list.head != nil {
		list.Delete(0)
		fmt.Println(list.Repr())
	}
}

