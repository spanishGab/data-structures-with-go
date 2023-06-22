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

  for !stack.isEmpty() {
    fmt.Println("Popping:", stack.Peek())
    stack.Pop()
    fmt.Println(stack.Repr())
  }
}

func Example2() {
  var stack Stack = New(4)

  stack.Push(1)
  fmt.Println(stack.Repr())
  fmt.Println("Popping:", stack.Pop())

  stack.Push(2)
  fmt.Println(stack.Repr())
  fmt.Println("Popping:", stack.Pop())

  stack.Push(4)
  fmt.Println(stack.Repr())
  fmt.Println("Popping:", stack.Pop())

  stack.Push(8)
  fmt.Println(stack.Repr())
  fmt.Println("Popping:", stack.Pop())
  fmt.Println(stack.Repr())
}
