package main

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	head     *Node
	inserted int
	tail     *Node
}

type Node struct {
	next  *Node
	value int
	prev  *Node
}

func Add(list *LinkedList, valor int) {
	node := &Node{next: nil, value: valor, prev: nil}

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node_atual := list.head

		for node_atual.next != nil {
			node_atual = node_atual.next
		}
		node.prev = node_atual
		node_atual.next = node
		list.tail = node

	}

	list.inserted++
}

func AddOnIndex(list *LinkedList, index int, valor int) error {
	if index < 0 || index > list.inserted {
		return errors.New("index out of lower or upper bound, try again")
	} else {
		i := 0
		node_novo := &Node{next: nil, value: valor, prev: nil} // no que vai ser adicionado

		node_atual := list.head
		var node_anterior *Node

		for i != index {
			node_anterior = node_atual
			node_atual = node_atual.next
			i++
		}

		if i == 0 {
			list.head = node_novo
			node_novo.next = node_atual
		} else {
			node_novo.next = node_atual
			node_anterior.next = node_novo
		}

		list.inserted++
	}

	return nil
}

func RemoveOnIndex(list *LinkedList, index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New("index out of lower or upper bound, try again")
	} else {
		i := 0
		node_atual := list.head
		var node_anterior *Node = nil

		for i != index {
			node_anterior = node_atual
			node_atual = node_atual.next
			i++
		}

		if i == 0 {
			list.head = list.head.next
		} else {
			node_anterior.next = node_atual.next
		}

		list.inserted--
	}

	return nil

}

func Get(list *LinkedList, index int) (int, error) {
	node_atual := list.head
	i := 0

	if index < 0 || index >= list.inserted {
		return -1, errors.New("index out of lower or upper bound, try again")
	}

	for i != index {
		node_atual = node_atual.next
		i++
	}

	return node_atual.value, nil
}

func Set(list *LinkedList, index int, valor int) error {
	if index < 0 || index >= list.inserted {
		return errors.New("index out of lower or upper bound, check again")
	}

	node_novo := &Node{next: nil, value: valor}
	i := 0

	if list.head == nil {
		list.head = node_novo
		list.tail = node_novo
		return nil
	} else {
		node_atual := list.head

		for i != index {
			node_atual = node_atual.next
			i++
		}

		node_atual.value = valor
	}

	return nil
}

func PrintList(list *LinkedList) {
	fmt.Println("---------------------------------------------")
	i := 0

	fmt.Println(list.head)

	for i < list.inserted {
		num, result := Get(list, i)

		// tratando erros
		if result != nil {
			fmt.Println(result)
			fmt.Println(num)
			return
		}

		fmt.Printf("Elemento %v da lista: %v\n", i, num)

		i++
	}
	fmt.Println(list.tail)
	fmt.Println("---------------------------------------------")
}

func main() {
	x := LinkedList{head: nil, inserted: 0}
	Add(&x, 10)
	Add(&x, 15)
	PrintList(&x)

	AddOnIndex(&x, 1, 20)
	Add(&x, 25)

	PrintList(&x)

	resultado_set := Set(&x, 2, 500)
	PrintList(&x)

	// tratando erros
	if resultado_set != nil {
		fmt.Println(resultado_set)
	}

	resultado_remove := RemoveOnIndex(&x, 1)

	if resultado_remove != nil {
		fmt.Println(resultado_remove)
	}

	PrintList(&x)

}
