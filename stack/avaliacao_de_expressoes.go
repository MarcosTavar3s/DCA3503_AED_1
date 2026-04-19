package main

import (
	"fmt"
)

func balparenteses(expressao string) bool {
	var avaliador Stack
	var state bool

	for _, r := range expressao {
		if r == '(' {
			avaliador.Push(1)
		} else if r == ')' {
			state = avaliador.IsEmpty()

			if state == false {
				avaliador.Pop()
			} else {
				return false
			}

		}
	}

	state = avaliador.IsEmpty()
	if state == false {
		return false
	}

	return true

}

func main() {
	var expressao string

	fmt.Println("Digite a expressao que sera avaliada: ")
	fmt.Scanln(&expressao)

	result := balparenteses(expressao)

	if result == false {
		fmt.Println("A expressao nao esta adequada")
	} else {
		fmt.Println("A expressao esta adequada")
	}

}
