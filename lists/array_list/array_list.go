package main

import (
	"errors"
	"fmt"
)

type ArrayList struct {
	vetor []int
	size  int
}

func (list *ArrayList) Init(val int) error {
	if val < 0 {
		return errors.New("value < 0")
	}

	list.vetor = make([]int, val)

	return nil
}

func (list *ArrayList) Add(val int) {
	if list.size == len(list.vetor) {
		aux := make([]int, 2*len(list.vetor))

		copy(aux, list.vetor)

		list.vetor = aux

	}

	list.vetor[list.size] = val
	list.size++

}

func (list *ArrayList) AddOnIndex(index int, val int) error {
	if index < 0 || index > list.size {
		return errors.New("index not available")
	}

	if index == len(list.vetor) {
		aux := make([]int, 2*len(list.vetor))
		copy(aux, list.vetor)
		list.vetor = aux
	}

	for i := list.size; i > index; i-- {
		list.vetor[i] = list.vetor[i-1]
	}

	list.vetor[index] = val
	list.size++

	return nil
}

func (list *ArrayList) RemoveOnIndex(index int) error {
	if index < 0 || index >= list.size {
		return errors.New("index not available")
	}

	for i := index; i < list.size; i++ {
		list.vetor[i] = list.vetor[i+1]
	}

	list.size--
	return nil
}

func (list *ArrayList) Get(index int) (int, error) {
	if index < 0 || index >= len(list.vetor) {
		return -1, errors.New("index not available")
	}

	return list.vetor[index], nil
}

func (list *ArrayList) Set(value int, index int) error {
	if index < 0 || index > list.size {
		return errors.New("index not available")
	}

	if index == len(list.vetor) {
		aux := make([]int, 2*len(list.vetor))
		copy(aux, list.vetor)
		list.vetor = aux
	}

	list.vetor[index] = value

	return nil

}

func (list *ArrayList) Size() int {
	return list.size
}

func (list *ArrayList) Reverse() {
	for i := 0; i < list.size/2; i++ {
		list.vetor[i], list.vetor[list.size-i] = list.vetor[list.size-i], list.vetor[i]
	}

}

func main() {
	lista := &ArrayList{}

	lista.Init(10)

	for i := 0; i < 12; i++ {
		lista.Add(i)
	}

	fmt.Println(lista.vetor)

	lista.RemoveOnIndex(0)
	fmt.Println(lista.vetor)

	lista.AddOnIndex(1, 55)
	fmt.Println(lista.vetor)

	fmt.Println(lista.Get(4))

	fmt.Println(lista.Set(60, 4))
	fmt.Println(lista.vetor)

	lista.Reverse()
	fmt.Println(lista.vetor)

}
