package main

import "fmt"

type QSort struct {
}

func (q *QSort) partition(arr []int, low, high int) int {
	i := low - 1
	pivot := high
	for j := low; j <= high-1; j++ {
		if arr[j] < arr[pivot] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[pivot] = arr[pivot], arr[i+1]
	return i + 1
}

func (q *QSort) qsort(arr []int, low, high int) {
	if low < high {
		pivot := q.partition(arr, low, high)
		q.qsort(arr, low, pivot-1)
		q.qsort(arr, pivot+1, high)
	}
}

func main() {
	arr := []int{10, 50, 30, 20, 80, 25}
	q := QSort{}
	q.qsort(arr, 0, len(arr)-1)

	for i := range arr {
		fmt.Print(arr[i], " ")
	}
}
