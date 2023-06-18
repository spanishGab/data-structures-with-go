package stacks

import "fmt"

func Example1() {
  var stack Stack = New(4)

  stack.Push(1)
  fmt.Println(stack.Repr())

  stack.Push(2)
  fmt.Println(stack.Repr())

  stack.Push(4)
  fmt.Println(stack.Repr())

  stack.Push(8)
  fmt.Println(stack.Repr())
}
