package main

import (
	"dsa/linkedLists/circular"
	"dsa/linkedLists/doubly"
	"dsa/linkedLists/singly"
	"dsa/stacks"
	"fmt"
)

func _singlyLinkedListExamples() {
	fmt.Println("Begin Example 1:")
	singly.Example1()
	fmt.Print("End Example 1\n\n")

	fmt.Println("Begin Example 2:")
	singly.Example2()
	fmt.Print("End Example 2\n\n")

	fmt.Println("Begin Example 3:")
	singly.Example3()
	fmt.Print("End Example 3\n\n")

	fmt.Println("Begin Example 4:")
	singly.Example4()
	fmt.Print("End Example 4\n\n")

	fmt.Println("Begin Example 5:")
	singly.Example5()
	fmt.Print("End Example 5\n\n")
}

func _doublyLinkedListExamples() {
	fmt.Println("Begin Example 1:")
	doubly.Example1()
	fmt.Print("End Example 1\n\n")

	fmt.Println("Begin Example 2:")
	doubly.Example2()
	fmt.Print("End Example 2\n\n")

	fmt.Println("Begin Example 3:")
	doubly.Example3()
	fmt.Print("End Example 3\n\n")

	fmt.Println("Begin Example 4:")
	doubly.Example4()
	fmt.Print("End Example 4\n\n")

	fmt.Println("Begin Example 5:")
	doubly.Example5()
	fmt.Print("End Example 5\n\n")
}

func _circularLinkedListExamples() {
	fmt.Println("Begin Example 1:")
	circular.Example1()
	fmt.Print("End Example 1\n\n")

	fmt.Println("Begin Example 2:")
	circular.Example2()
	fmt.Print("End Example 2\n\n")

	fmt.Println("Begin Example 3:")
	circular.Example3()
	fmt.Print("End Example 3\n\n")

	fmt.Println("Begin Example 4:")
	circular.Example4()
	fmt.Print("End Example 4\n\n")

	fmt.Println("Begin Example 5:")
	circular.Example5()
	fmt.Print("End Example 5\n\n")
}

func _stackExamples() {
	fmt.Println("Begin Example 1:")
	stacks.Example1()
	fmt.Print("End Example 1\n\n")

	fmt.Println("Begin Example 2:")
	stacks.Example2()
	fmt.Print("End Example 2\n\n")

	// fmt.Println("Begin Example 3:")
	// circular.Example3()
	// fmt.Print("End Example 3\n\n")

	// fmt.Println("Begin Example 4:")
	// circular.Example4()
	// fmt.Print("End Example 4\n\n")

	// fmt.Println("Begin Example 5:")
	// circular.Example5()
	// fmt.Print("End Example 5\n\n")
}

func main() {
	// use this space to execute some of the examples above
	// _singlyLinkedListExamples()
  // _doublyLinkedListExamples()
  // _circularLinkedListExamples()
  _stackExamples()
}
