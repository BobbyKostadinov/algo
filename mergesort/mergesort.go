package mergesort

import (
	"fmt"
	"math"
)

// MergeSort entry point
func MergeSort(A []int) []int {
	if len(A) <= 1 {
		return A
	}
	return sort(A)
}

func sort(A []int) []int {
	// Baser case
	if len(A) <= 1 {
		fmt.Println("Reached base case")
		return A
	}
	// define mid point
	q := int(math.Floor(float64(len(A)-1) / 2))

	fmt.Println("Passed Array ", A)
	fmt.Println("Mid Point ", q)

	A1 := sort(A[:q+1])
	A2 := sort(A[q+1:])

	mergeA := merge(A1, A2)
	fmt.Println("Merged: ", mergeA)
	return mergeA
}

func merge(A1 []int, A2 []int) []int {
	merged := make([]int, 0, (len(A1) + len(A2)))

	fmt.Println("Compare array A1 and A2", A1, A2)

	for i := 0; i <= cap(merged); i++ {
		if len(A1) < 1 {
			merged = append(merged, A2...)
			break
		}
		if len(A2) < 1 {
			merged = append(merged, A1...)
			break
		}
		if A1[0] < A2[0] {
			merged = append(merged, A1[0])
			A1 = A1[1:]
		} else {
			merged = append(merged, A2[0])
			A2 = A2[1:]
		}
	}

	return merged
}
