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

func Push(stack *Stack, val int) {
	no := &Node{next: stack.top, val: val}
	stack.top = no
	stack.inserted++
	fmt.Printf("%v inserido no topo da pilha, tamanho da pilha: %v\n", val, stack.inserted)
}

func Pop(stack *Stack) (int, error) {
	if stack.inserted == 0 {
		return -1, errors.New("A lista esta vazia")
	}

	value := stack.top.val
	stack.top = stack.top.next
	stack.inserted--

	fmt.Printf("%v removido do topo da pilha com o metodo pop, novo tamanho da pilha: %v\n", value, stack.inserted)
	return value, nil
}

func IsEmpty(stack *Stack) bool {
	if stack.inserted == 0 {
		fmt.Printf("A lista esta vazia\n")
		return true
	}

	fmt.Printf("A lista tem %v elemento(s)\n", stack.inserted)
	return false
}

func Top(stack *Stack) (int, error) {
	if stack.inserted == 0 {
		return -1, errors.New("A lista esta vazia")
	}

	fmt.Printf("Topo da lista com metodo top: %v\n", stack.top.val)

	return stack.top.val, nil
}

func ShowAllItems(stack *Stack) {
	aux := stack.top

	for i := 0; i < stack.inserted; i++ {
		fmt.Printf("%v na posicao %v\n", aux.val, i)
		aux = aux.next
	}
}

func main() {
	stack := Stack{top: nil, inserted: 0}

	IsEmpty(&stack)
	Push(&stack, 10)
	Push(&stack, 100)
	Push(&stack, 200)
	ShowAllItems(&stack)

	Pop(&stack)

	Top(&stack)
	IsEmpty(&stack)
}
