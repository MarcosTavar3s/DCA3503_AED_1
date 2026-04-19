package main

import "errors"

type LinkedListQueue struct {
	front *Node
	rear  *Node
	size  int
}

type Node struct {
	val  int
	prev *Node
	next *Node
}

func (queue *LinkedListQueue) Enqueue(val int) {
	newNode := &Node{val: val, next: nil}

	if queue.size == 0 {
		queue.front = newNode
	} else {
		queue.rear.next = newNode
	}

	queue.rear = newNode
	queue.size++
}

func (queue *LinkedListQueue) Dequeue() (int, error) {
	var aux *Node

	if queue.size == 0 {
		return -1, errors.New("empty")
	}

	aux = queue.front

	if queue.size == 1 {
		queue.rear = nil
		queue.front = nil
	}
	queue.front = aux.next
	queue.size--

	return aux.val, nil
}

func (queue *LinkedListQueue) Front() (int, error) {
	if queue.size == 0 {
		return -1, errors.New("error msg")
	}

	return queue.front.val, nil
}

func (queue *LinkedListQueue) IsEmpty() bool {
	if queue.size == 0 {
		return true
	}
	return false
}

func (queue *LinkedListQueue) Size() int {
	return queue.size
}
