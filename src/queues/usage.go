package queues

import "fmt"

func Example1() {
  var queue CircularQueue = New(4)

  queue.Enqueue(100)
  fmt.Println(queue.Repr())

  queue.Enqueue(101)
  fmt.Println(queue.Repr())

  queue.Enqueue(102)
  fmt.Println(queue.Repr())

  queue.Enqueue(103)
  fmt.Println(queue.Repr())

  fmt.Println("Peek:", queue.Peek())

  for !queue.isEmpty() {
    fmt.Println("Popping", queue.Dequeue())
    fmt.Println(queue.Repr())
  }
}

func Example2() {
  var queue CircularQueue = New(4)

  for i := 0; i < 4; i++ {
    queue.Enqueue(100 + i)
    fmt.Println(queue.Repr())
  }
  queue.Dequeue()
  fmt.Println(queue.Repr())

  queue.Dequeue()
  fmt.Println(queue.Repr())

  queue.Enqueue(200)
  fmt.Println(queue.Repr())

  queue.Enqueue(201)
  fmt.Println(queue.Repr())
}
