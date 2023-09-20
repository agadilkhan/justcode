package algorithms

import "composition-test/sort"

type Quick struct {
	Arr []int
}

func NewQuick(arr []int) *Quick {
	return &Quick{
		Arr: arr,
	}
}

func (q *Quick) IsMatch(algorithmType sort.AlorithmType) bool {
	return "quick" == algorithmType
}

func (q *Quick) Sort() {
	q.qsort(q.Arr, 0, len(q.Arr)-1)
}

func (q *Quick) partition(arr []int, low, high int) int {
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

func (q *Quick) qsort(arr []int, low, high int) {
	if low < high {
		pivot := q.partition(arr, low, high)
		q.qsort(arr, low, pivot-1)
		q.qsort(arr, pivot+1, high)
	}
}
