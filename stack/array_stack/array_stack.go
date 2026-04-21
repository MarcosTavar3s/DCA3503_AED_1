package main

import (
	"errors"
	"fmt"
)

type ArrayStack struct {
	vetor    []int
	inserted int
}

func (stack *ArrayStack) Init(val int) error {
	if val < 0 {
		return errors.New("init error")
	}

	stack.vetor = make([]int, val)

	return nil
}

func (stack *ArrayStack) Push(val int) {
	if stack.inserted == len(stack.vetor) {
		aux := make([]int, 2*len(stack.vetor))
		copy(aux, stack.vetor)
		stack.vetor = aux
	}

	stack.vetor[stack.inserted] = val
	stack.inserted++
}

func (stack *ArrayStack) Pop() (int, error) {
	if stack.inserted == 0 {
		return -1, errors.New("empty stack")
	}

	aux := stack.vetor[stack.inserted-1]
	stack.inserted--
	return aux, nil
}

func (stack *ArrayStack) Peek() (int, error) {
	if stack.inserted == 0 {
		return -1, errors.New("empty stack")
	}

	aux := stack.vetor[stack.inserted-1]
	return aux, nil
}

func (stack *ArrayStack) IsEmpty() bool {
	if stack.inserted == 0 {
		return true
	}

	return false

}

func (stack *ArrayStack) Size() int {
	return stack.inserted
}

func main() {
	stack := &ArrayStack{}

	stack.Init(5)

	for i := 0; i < 11; i++ {
		stack.Push((i + 1) * 10)
	}

	fmt.Println("-----")
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Peek())

	fmt.Println("-----")
	fmt.Println(stack.vetor, stack.inserted)
	for i := 0; i < 10; i++ {
		fmt.Println(stack.Pop())
	}

	fmt.Println("-----")
	fmt.Println(stack.Peek())

}
