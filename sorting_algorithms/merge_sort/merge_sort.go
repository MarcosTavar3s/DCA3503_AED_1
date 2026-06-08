package main

import (
	"fmt"
	"math/rand/v2"
)

func merge(v []int, d []int, e []int) {
	iv := 0
	id := 0
	ie := 0

	for id < len(d) && ie < len(e) {
		if d[id] > e[ie] {
			v[iv] = e[ie]
			ie++

		} else {
			v[iv] = d[id]
			id++
		}
		iv++
	}

	for id < len(d) {
		v[iv] = d[id]
		id++
		iv++
	}

	for ie < len(e) {
		v[iv] = e[ie]
		ie++
		iv++
	}
}

func MergeSort(v []int) {
	if len(v) > 1 {
		tamE := len(v) / 2
		tamD := len(v) - tamE

		e := make([]int, tamE)
		d := make([]int, tamD)

		for i := 0; i < tamE; i++ {
			e[i] = v[i]
		}

		for i := 0; i < tamD; i++ {
			d[i] = v[tamE+i]
		}

		MergeSort(e)
		MergeSort(d)
		merge(v, d, e)
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
	fmt.Println("Merge sort algorithm")

	size := 50
	v := generateVector(size)
	printVector(v)
	MergeSort(v)
	printVector(v)
}
