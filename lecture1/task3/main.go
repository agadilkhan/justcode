package main

import (
	"fmt"
	"sort"
)

func CompareTwoSlices(arr, arr1 []int64) bool {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})

	if len(arr1) != len(arr) {
		return false
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != arr1[i] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int64{1, 2, 3}
	arr1 := []int64{3, 1, 2}

	fmt.Print(CompareTwoSlices(arr, arr1))
}
