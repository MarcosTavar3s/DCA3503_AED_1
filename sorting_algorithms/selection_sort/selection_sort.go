package main

import (
	"fmt"
	"math/rand/v2"
)

func selectionSort(v []int) {

	for i := 0; i < len(v)-1; i++ {
		iMenor := i

		for j := i + 1; j < len(v); j++ {
			if v[j] < v[iMenor] {
				iMenor = j
			}
		}

		v[i], v[iMenor] = v[iMenor], v[i]

	}
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
	fmt.Println("Selection sort algorithm")

	size := 50
	v := generateVector(size)
	printVector(v)
	selectionSort(v)
	printVector(v)
}
