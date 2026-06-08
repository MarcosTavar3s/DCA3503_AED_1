package main

import (
	"fmt"
	"math/rand/v2"
)

func countingSort(v []int) []int {
	menor := v[0]
	maior := v[0]

	for i := 1; i < len(v); i++ {
		if menor > v[i] {
			menor = v[i]
		}

		if maior < v[i] {
			maior = v[i]
		}
	}

	tamC := maior - menor + 1
	c := make([]int, tamC)

	for i := 0; i < len(v); i++ {
		index := v[i] - menor
		c[index]++
	}

	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	novoVetor := make([]int, len(v))

	for i := 0; i < len(novoVetor); i++ {
		indexOrd := c[v[i]-menor] - 1
		novoVetor[indexOrd] = v[i]
		c[v[i]-menor]--
	}

	return novoVetor
}

func generateVector(size int) []int {
	v := make([]int, size)

	for i := 0; i < size; i++ {
		v[i] = rand.IntN(100)
	}

	return v
}

func printVector(v []int) {
	for i := 0; i < len(v); i++ {
		if i == 0 {
			fmt.Printf("[ %v", v[i])
		} else if i == len(v)-1 {
			fmt.Printf("%v ]\n", v[i])
		} else {
			fmt.Printf(" %v ", v[i])
		}

	}
}

func main() {
	fmt.Println("Merge sort algorithm")

	size := 50
	v := generateVector(size)
	printVector(v)
	v = countingSort(v)
	printVector(v)
}
