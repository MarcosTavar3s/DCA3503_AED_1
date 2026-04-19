package main

import "fmt"

func Array() {
	fmt.Println("Array List Queue")
	array_fifo := &ArrayQueue{}

	array_fifo.Init(2)
	array_fifo.Enqueue(10)
	array_fifo.Enqueue(20)
	fmt.Println(array_fifo.Dequeue())

	array_fifo.Enqueue(40)
	fmt.Println(array_fifo.v)

	array_fifo.Enqueue(60)
	fmt.Println(array_fifo.v)
	fmt.Println(array_fifo.front)
	fmt.Println(array_fifo.back)

	fmt.Println(array_fifo.Dequeue())
	fmt.Println(array_fifo.v)
	fmt.Println(array_fifo.front)
	fmt.Println(array_fifo.back)
}

func Linked() {
	fmt.Println("Linked List Queue")
	linked_fifo := &LinkedListQueue{}

	linked_fifo.Enqueue(10)
	linked_fifo.Enqueue(20)
	fmt.Println(linked_fifo.Dequeue())

	linked_fifo.Enqueue(40)

	linked_fifo.Enqueue(60)
	fmt.Println(linked_fifo.front.val)
	fmt.Println(linked_fifo.rear.val)

	fmt.Println(linked_fifo.Dequeue())
	fmt.Println(linked_fifo.front.val)
	fmt.Println(linked_fifo.rear.val)
}

func main() {
	Array()
	Linked()
}
