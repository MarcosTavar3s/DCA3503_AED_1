package main

import "fmt"

func busca_linear(vetor []int, val int) int {
	for i := 0; i < len(vetor); i++ {
		if vetor[i] == val {
			return i
		}
	}
	return -1
}

func busca_binc_recursiva(vetor []int, inicio int, fim int, val int) int {

	if inicio > fim {
		return -1
	}

	meio := (inicio + fim) / 2
	if vetor[meio] == val {
		return meio
	}

	if vetor[meio] > val {
		return busca_binc_recursiva(vetor, inicio, meio-1, val)
	}

	return busca_binc_recursiva(vetor, meio+1, fim, val)

}

func busca_binc_iterativa(vetor []int, inicio int, fim int, val int) int {

	for true {
		meio := (inicio + fim) / 2

		if inicio > fim {
			break
		}

		if vetor[meio] == val {
			return meio
		}

		if vetor[meio] > val {
			fim = meio - 1
		} else {
			inicio = meio + 1
		}

	}

	return -1

}

func busca_bind_recursiva(vetor []int, inicio int, fim int, val int) int {

	if inicio > fim {
		return -1
	}

	meio := (inicio + fim) / 2
	if vetor[meio] == val {
		return meio
	}

	if vetor[meio] > val {
		return busca_binc_recursiva(vetor, meio+1, fim, val)
	}

	return busca_bind_recursiva(vetor, inicio, meio-1, val)
}

func busca_bind_iterativa(vetor []int, inicio int, fim int, val int) int {

	for true {
		meio := (inicio + fim) / 2

		if inicio > fim {
			break
		}

		if vetor[meio] == val {
			return meio
		}

		if vetor[meio] > val {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}

	}

	return -1
}

func main() {
	vetor := make([]int, 15)

	for i := 0; i < len(vetor); i++ {
		vetor[i] = i*4 + 7
	}

	fmt.Println("------------------")
	fmt.Print("Vetor crescente: ")
	fmt.Println(vetor)

	fmt.Print("Busca linear: ")
	fmt.Println(busca_linear(vetor, 19))

	fmt.Print("Busca binaria crescente recursiva: ")
	fmt.Println(busca_binc_recursiva(vetor, 0, len(vetor), 19))

	fmt.Print("Busca binaria crescente iterativa: ")
	fmt.Println(busca_binc_iterativa(vetor, 0, len(vetor), 63))

	for i := 0; i < len(vetor); i++ {
		vetor[i] = i*(-4) - 7
	}

	fmt.Println("------------------")
	fmt.Print("Vetor decrescente: ")
	fmt.Println(vetor)

	fmt.Print("Busca linear: ")
	fmt.Println(busca_linear(vetor, -27))

	fmt.Print("Busca binaria decrescente recursiva: ")
	fmt.Println(busca_bind_recursiva(vetor, 0, len(vetor), -51))

	fmt.Print("Busca binaria decrescente iterativa: ")
	fmt.Println(busca_bind_iterativa(vetor, 0, len(vetor), -51))
}
