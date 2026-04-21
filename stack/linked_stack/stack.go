package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	top      *Node
	inserted int
}

type Node struct {
	next *Node
	val  int
}

func (stack *Stack) Push(val int) {
	no := &Node{next: stack.top, val: val}
	stack.top = no
	stack.inserted++
	// fmt.Printf("%v inserido no topo da pilha, tamanho da pilha: %v\n", val, stack.inserted)
}

func (stack *Stack) Pop() (int, error) {
	if stack.inserted == 0 {
		return -1, errors.New("A lista esta vazia")
	}

	value := stack.top.val
	stack.top = stack.top.next
	stack.inserted--

	// fmt.Printf("%v removido do topo da pilha com o metodo pop, novo tamanho da pilha: %v\n", value, stack.inserted)
	return value, nil
}

func (stack *Stack) IsEmpty() bool {
	if stack.inserted == 0 {
		// fmt.Printf("A lista esta vazia\n")
		return true
	}

	// fmt.Printf("A lista tem %v elemento(s)\n", stack.inserted)
	return false
}

func (stack *Stack) Top() (int, error) {
	if stack.inserted == 0 {
		return -1, errors.New("A lista esta vazia")
	}

	// fmt.Printf("Topo da lista com metodo top: %v\n", stack.top.val)

	return stack.top.val, nil
}

func (stack *Stack) ShowAllItems() {
	aux := stack.top

	for i := 0; i < stack.inserted; i++ {
		fmt.Printf("%v na posicao %v\n", aux.val, i)
		aux = aux.next
	}
}

func stack_demo() {
	stack := Stack{top: nil, inserted: 0}

	stack.IsEmpty()
	stack.Push(10)
	stack.Push(100)
	stack.Push(200)
	stack.ShowAllItems()

	stack.Pop()

	stack.Top()
	stack.IsEmpty()
}
