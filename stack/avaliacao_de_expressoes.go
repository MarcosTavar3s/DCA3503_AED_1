package main

import (
	"fmt"
	"os"
)

func main() {
	var avaliador Stack
	var expressao string

	fmt.Println("Digite a expressao que sera avaliada: ")
	fmt.Scanln(&expressao)

	state := true

	for _, r := range expressao {
		if r == '(' {
			avaliador.Push(1)
		} else if r == ')' {
			state = avaliador.IsEmpty()

			if state == false {
				avaliador.Pop()
			} else {
				fmt.Println("A expressao nao esta adequada")
				os.Exit(0)
			}

		}
	}

	state = avaliador.IsEmpty()

	if state == false {
		fmt.Println("A expressao nao esta adequada")
	} else {
		fmt.Println("A expressao esta adequada")
	}

}
