package main

import (
	"fmt"

	"github.com/BobbyKostadinov/algo/sort"
)

func main() {
	A := []int{1, 4, 5, 3, 23, 2, -3, 334, 51, 54, 31, 13, 11, 23, 4, -34, 100, 44, 23, 13, 7, 125, 0, 2, 32, 41, -1}
	fmt.Println("Original Arr")
	fmt.Printf("Selection results: %v \n", sort.Selection(A))

}
