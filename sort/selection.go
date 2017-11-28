package sort

// Selection sort method - On2
func Selection(A []int) []int {
	for i := range A {
		minIndex := i
		j := i
		for j < len(A) {
			if A[minIndex] > A[j] {
				minIndex = j

			}
			j++
		}
		if swapped := A[i]; A[minIndex] < A[i] {
			A[i] = A[minIndex]
			A[minIndex] = swapped

		}
	}
	return A
}
