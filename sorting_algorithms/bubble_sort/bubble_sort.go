package main

import (
	"fmt"
	"math/rand/v2"
)

func bubbleSort(v []int) {

	for i := 0; i < len(v)-1; i++ {
		mudou := false

		for j := 0; j < len(v)-i-1; j++ {
			if v[j] > v[j+1] {
				v[j+1], v[j] = v[j], v[j+1]
				mudou = true
			}
		}
		if !mudou {
			break
		}
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
	fmt.Println("Bubble sort algorithm")

	size := 50
	v := generateVector(size)
	printVector(v)
	bubbleSort(v)
	printVector(v)
}
