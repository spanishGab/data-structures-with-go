package queues

import (
	"strconv"
)

type CircularQueue struct {
	capacity int
	rear     int
	front    int
	data     []int
}

func New(capacity int) CircularQueue {
  data := make([]int, capacity)
	return CircularQueue{
		capacity: capacity,
		rear:     -1,
		front:    -1,
		data:     data,
	}
}

func (queue CircularQueue) isFull() bool {
	return (queue.front == 0 && queue.rear == queue.capacity-1) || queue.front == queue.rear+1
}

func (queue *CircularQueue) Enqueue(value int) {
  if queue.isFull() {
    panic("Queue is full! Cannot add more elements to it")
  }

  if queue.front == -1 {
    queue.front = 0
  }
  queue.rear = (queue.rear + 1) % queue.capacity
  queue.data[queue.rear] = value
}

func (queue *CircularQueue) isEmpty() bool {
  return queue.front == -1
}

func (queue *CircularQueue) Dequeue() int {
  if queue.isEmpty() {
    panic("Queue is empty! Cannot remove any element from it")
  }

  element := queue.data[queue.front]
  if queue.front == queue.rear {
    queue.front = -1
    queue.rear = -1
  } else {
    queue.front = (queue.front + 1) % queue.capacity
  }
  return element
}

func (queue *CircularQueue) Peek() int {
  if queue.isEmpty() {
    panic("Queue is empty!")
  }
  return queue.data[queue.front]
}

func (queue *CircularQueue) Repr() string {
  repr := "Queue content: "
  if queue.isEmpty() {
    return repr + "-|"
  }
  i := queue.front
  for ; i != queue.rear; i = (i + 1) % queue.capacity {
    repr += "-|" + strconv.FormatInt(int64(queue.data[i]), 10)
  }
  repr += "-|" + strconv.FormatInt(int64(queue.data[i]), 10)
  return repr
}
