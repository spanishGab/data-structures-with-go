package stacks

import (
	"dsa/linkedLists/circular"
	"strings"
)

type Stack struct {
  top int
  capacity int
  data circular.CircularLinkedList
}

func New(capacity int) Stack {
  var data circular.CircularLinkedList = circular.New()
  return Stack{
    top: -1,
    capacity: capacity,
    data: data,
  }
}

func (stack *Stack) isFull() bool {
  return stack.top == stack.capacity - 1
}

func (stack *Stack) Push(value int) {
  if (stack.isFull()) {
    panic("Stack is full! Cannot add more elements to it")
  }

  stack.top++
  stack.data.Append(value)
}

func (stack *Stack) Repr() string {
  var stackElements []string
  stackElements = strings.Split(stack.data.Repr(), "<->")
  stackElements = stackElements[1:]
  stackElements = stackElements[:len(stackElements)-1]

  stackLength := len(stackElements)
  repr := "Stack content:\n"
  repr += stackElements[stackLength-1] + " <-TOP\n"
  for i := stackLength-2; i >= 0; i-- {
    repr += stackElements[i] + "\n"
  }
  return repr[:len(repr)-1]
}
