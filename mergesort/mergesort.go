package mergesort

import "math"

// MergeSort entry point
func Sort(A []int) []int {
	if len(A) <= 1 {
		return A
	}
	return sortRec(A)
}

func sortRec(A []int) []int {
	// Base case
	if len(A) <= 1 {
		return A
	}
	// define mid point
	q := int(math.Floor(float64(len(A)-1) / 2))

	A1 := sortRec(A[:q+1])
	A2 := sortRec(A[q+1:])

	mergeA := merge(A1, A2)
	return mergeA
}

func merge(A1 []int, A2 []int) []int {
	merged := make([]int, 0, (len(A1) + len(A2)))

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
