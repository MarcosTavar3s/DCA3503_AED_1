package main

import (
	"errors"
)

type ArrayQueue struct {
	v     []int
	front int
	back  int
	size  int
}

func (queue *ArrayQueue) Init(size int) {
	queue.v = make([]int, size)
	queue.front = 0
	queue.back = -1
}

func (queue *ArrayQueue) Enqueue(val int) {
	// se a fila estiver lotada -> duplicar
	if queue.size == len(queue.v) {
		aux := make([]int, 2*queue.size)

		count := 0
		for i := queue.front; count < queue.size; i = (i + 1) % len(queue.v) {
			aux[count] = queue.v[i]
			count++
		}

		queue.front = 0
		queue.v = aux
		queue.v[queue.size] = val
		queue.back = queue.size

	} else {
		queue.back = (queue.back + 1) % len(queue.v)
		queue.v[queue.back] = val

	}
	queue.size++

}

func (queue *ArrayQueue) Dequeue() (int, error) {
	if queue.size == 0 {
		return -1, errors.New("empty queue")
	}

	aux := queue.v[queue.front]
	queue.front = (queue.front + 1) % len(queue.v)

	queue.size--
	return aux, nil

}

func (queue *ArrayQueue) Front() (int, error) {
	if queue.size == 0 {
		return -1, errors.New("error msg")
	}

	return queue.v[queue.front], nil

}

func (queue *ArrayQueue) IsEmpty() bool {
	if queue.size == 0 {
		return true
	}
	return false

}

func (queue *ArrayQueue) Size() int {
	return queue.size
}
