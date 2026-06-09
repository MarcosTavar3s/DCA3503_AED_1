package main

import (
	"fmt"
	"math/rand/v2"
)

func QuickSort(v []int, ini int, fim int) {
	if ini < fim {
		indexPivot := Partition(v, ini, fim)
		QuickSort(v, ini, indexPivot-1)
		QuickSort(v, indexPivot+1, fim)
	}
}

func Partition(v []int, ini int, fim int) int {
	rand := rand.IntN(fim-ini+1) + ini

	v[fim], v[rand] = v[rand], v[fim] // faz o pivo ser aleatorio

	pivot := v[fim]
	indexPivot := ini

	for i := ini; i < fim; i++ {
		if v[i] <= pivot {
			v[indexPivot], v[i] = v[i], v[indexPivot]
			indexPivot++
		}
	}

	v[fim], v[indexPivot] = v[indexPivot], v[fim]

	return indexPivot
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
	fmt.Println("Quick sort algorithm")

	size := 50
	v := generateVector(size)
	printVector(v)
	QuickSort(v, 0, len(v)-1)
	printVector(v)
}
