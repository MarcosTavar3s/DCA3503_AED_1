package main

import (
	"errors"
	"fmt"
)

type DoubleLinkedList struct {
	head     *Node
	inserted int
	tail     *Node
}

type Node struct {
	prev  *Node
	value int
	next  *Node
}

func Add(list *DoubleLinkedList, valor int) {
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

func AddOnIndex(list *DoubleLinkedList, index int, valor int) error {
	if index < 0 || index > list.inserted {
		return errors.New("index out of lower or upper bound, try again")
	} else {
		i := 0
		node_novo := &Node{next: nil, value: valor, prev: nil} // no que vai ser adicionado

		switch index {
		case 0:
			if list.inserted == 0 {
				list.head = node_novo
				list.tail = node_novo
			} else {
				node_novo.next = list.head
				list.head.prev = node_novo
				list.head = node_novo
			}

		case list.inserted:
			node_novo.prev = list.tail
			list.tail.next = node_novo
			list.tail = node_novo

		default:
			node_atual := list.head

			for i != index {
				node_atual = node_atual.next
				i++
			}

			node_novo.next = node_atual.next
			node_atual.next = node_novo
			node_novo.prev = node_atual
		}

		list.inserted++
	}

	return nil
}

func RemoveOnIndex(list *DoubleLinkedList, index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New("index out of lower or upper bound, try again")
	} else {
		switch index {
		case 0:
			if list.inserted == 1 {
				list.head = nil
				list.tail = nil
			} else {
				next := list.head.next
				next.prev = nil
				list.head = next
			}
		case list.inserted - 1:
			prev := list.tail.prev
			prev.next = nil
			list.tail = prev

		default:
			node_atual := list.head
			i := 0

			for i != index {
				node_atual = node_atual.next
				i++
			}

			node_atual.prev.next = node_atual.next
			node_atual.next.prev = node_atual.prev

		}
	}

	list.inserted--

	return nil

}

func Get(list *DoubleLinkedList, index int) (int, error) {
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

func Set(list *DoubleLinkedList, index int, valor int) error {
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

func PrintList(list *DoubleLinkedList) {
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

func Reverse(list *DoubleLinkedList) {
	if list.inserted == 0 {
		fmt.Println("Lista vazia")
	} else {
		if list.inserted == 1 {
			fmt.Println("Lista so tem 1 elemento")
		} else {
			aux := list.head

			for i := 0; i < list.inserted; i++ {

				aux.prev, aux.next = aux.next, aux.prev
				fmt.Println(aux)
				aux = aux.prev

			}

			list.tail, list.head = list.head, list.tail

		}

	}

}

func main() {
	x := DoubleLinkedList{head: nil, inserted: 0}
	Add(&x, 10)
	Add(&x, 15)
	Add(&x, 20)
	Add(&x, 25)

	PrintList(&x)

	AddOnIndex(&x, 2, 20)
	Add(&x, 25)

	PrintList(&x)

	resultado_set := Set(&x, 2, 500)
	PrintList(&x)

	// tratando erros
	if resultado_set != nil {
		fmt.Println(resultado_set)
	}

	resultado_remove := RemoveOnIndex(&x, 3)
	resultado_remove = RemoveOnIndex(&x, 2)
	resultado_remove = RemoveOnIndex(&x, 1)

	if resultado_remove != nil {
		fmt.Println(resultado_remove)
	}
	fmt.Println(x.inserted)

	aux := x.head
	for i := 0; i < x.inserted; i++ {
		fmt.Println(aux)
		aux = aux.next

	}
	fmt.Println("-----------------")

	Reverse(&x)
	PrintList(&x)

}
