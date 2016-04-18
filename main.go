package main

import "github.com/BobbyKostadinov/algo/mergesort"

func main() {
	A := []int{1, 4, 5, 3, 23, 2, 54, 31, 13, 0, 2, 32, 41, -1}
	mergesort.MergeSort(A)
}
