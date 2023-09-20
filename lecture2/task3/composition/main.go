package main

import (
	"composition-test/sort"
	"composition-test/sort/algorithms"
	"fmt"
)

func main() {
	var (
		quickSort  = algorithms.Quick{[]int{9, 10, 8, 4, 5}}
		bubbleSort = algorithms.Bubble{[]int{20, 30, 1, 3, 2, 19}}

		sortingManager = sort.NewManager(&quickSort, &bubbleSort)
	)

	sortingManager.Execute(sort.Quick)
	sortingManager.Execute(sort.Bubble)

	fmt.Println(quickSort)
	fmt.Println(bubbleSort)
}
