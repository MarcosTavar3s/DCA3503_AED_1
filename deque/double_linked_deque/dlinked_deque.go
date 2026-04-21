package main

import "errors"

type DoubleLinkedDeque struct {
	front    *Node
	inserted int
	rear     *Node
}

type Node struct {
	prev *Node
	val  int
	next *Node
}

func (deque *DoubleLinkedDeque) EnqueueFront(value int) {
	newNode := &Node{val: value}

	if deque.inserted == 0 {
		deque.rear = newNode
	} else {
		newNode.next = deque.front
		deque.front.prev = newNode
	}

	deque.front = newNode
	deque.inserted++

}

func (deque *DoubleLinkedDeque) EnqueueRear(value int) {
	newNode := &Node{val: value}

	if deque.inserted == 0 {
		deque.front = newNode
	} else {
		newNode.prev = deque.rear
		deque.rear.next = newNode
	}

	deque.rear = newNode
	deque.inserted++

}

func (deque *DoubleLinkedDeque) DequeueFront() (int, error) {
	if deque.inserted == 0 {
		return -1, nil
	}

	aux := deque.front.val

	if deque.inserted == 1 {
		deque.front = nil
		deque.rear = nil
	} else {
		deque.front = deque.front.next
		deque.front.prev = nil

	}
	deque.inserted--

	return aux, nil
}

func (deque *DoubleLinkedDeque) DequeueRear() (int, error) {
	if deque.inserted == 0 {
		return -1, nil
	}

	aux := deque.rear.val

	if deque.inserted == 1 {
		deque.front = nil
		deque.rear = nil
	} else {
		deque.rear = deque.rear.prev
		deque.rear.next = nil
	}

	deque.inserted--

	return aux, nil

}

func (deque *DoubleLinkedDeque) Front() (int, error) {
	if deque.inserted == 0 {
		return -1, errors.New("empty deque")
	}
	return deque.front.val, nil
}

func (deque *DoubleLinkedDeque) Rear() (int, error) {
	if deque.inserted == 0 {
		return -1, errors.New("empty deque")
	}
	return deque.rear.val, nil
}

func (deque *DoubleLinkedDeque) IsEmpty() bool {
	if deque.inserted == 0 {
		return true
	}
	return false
}

func (deque *DoubleLinkedDeque) Size() int {
	return deque.inserted
}

func main() {

}
