package algorithms

import "composition-test/sort"

type Bubble struct {
	Arr []int
}

func NewBubble(arr []int) *Bubble {
	return &Bubble{arr}
}

func (b *Bubble) IsMatch(alorithmType sort.AlorithmType) bool {
	return "bubble" == alorithmType
}

func (b *Bubble) Sort() {
	n := len(b.Arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if b.Arr[j] > b.Arr[j+1] {
				b.Arr[j], b.Arr[j+1] = b.Arr[j+1], b.Arr[j]
			}
		}
	}
}
