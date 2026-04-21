package main

import "errors"

type ArrayDeque struct {
	inserted int
	vetor    []int
	front    int
	rear     int
}

func (deque *ArrayDeque) Init(val int) error {
	if val <= 0 {
		return errors.New("init error")
	}

	deque.vetor = make([]int, val)
	return nil
}

func (deque *ArrayDeque) EnqueueFront(val int) {
	if deque.inserted == len(deque.vetor) {
		aux := make([]int, 2*len(deque.vetor))
		c := 0

		for i := deque.front; c < deque.inserted; i = (i + 1) % deque.inserted {
			aux[c] = deque.vetor[i]
			c++
		}

		deque.vetor = aux
		deque.front = 0
		deque.rear = deque.inserted
	}

	deque.front = (deque.front - 1 + len(deque.vetor)) % len(deque.vetor)
	deque.vetor[deque.front] = val
	deque.inserted++
}

func (deque *ArrayDeque) EnqueueRear(val int) {
	if deque.inserted == len(deque.vetor) {
		aux := make([]int, 2*len(deque.vetor))
		c := 0

		for i := deque.front; c < deque.inserted; i = (i + 1) % deque.inserted {
			aux[c] = deque.vetor[i]
			c++
		}
		deque.vetor = aux
		deque.front = 0
		deque.rear = deque.inserted

	}

	deque.vetor[deque.rear] = val
	deque.rear = (deque.rear + 1) % len(deque.vetor)
	deque.inserted++
}

func (deque *ArrayDeque) DequeueFront() (int, error) {
	if deque.inserted == 0 {
		return -1, errors.New("empty deque")
	}

	aux := deque.vetor[deque.front]
	deque.front = (deque.front + 1) % len(deque.vetor)

	deque.inserted--
	return aux, nil
}

func (deque *ArrayDeque) DequeueRear() (int, error) {
	if deque.inserted == 0 {
		return -1, errors.New("empty deque")
	}

	deque.rear = (deque.rear - 1 + len(deque.vetor)) % len(deque.vetor)
	aux := deque.vetor[deque.rear]

	deque.inserted--
	return aux, nil
}

func (deque *ArrayDeque) Front() (int, error) {
	if deque.inserted == 0 {
		return -1, errors.New("empty deque")
	}

	return deque.vetor[deque.front], nil
}

func (deque *ArrayDeque) Rear() (int, error) {
	if deque.inserted == 0 {
		return -1, errors.New("empty deque")
	}

	return deque.vetor[(deque.rear-1+len(deque.vetor))%len(deque.vetor)], nil
}

func (deque *ArrayDeque) IsEmpty() bool {
	if deque.inserted == 0 {
		return true
	}

	return false
}

func (deque *ArrayDeque) Size() int {
	return deque.inserted
}

func main() {

}
