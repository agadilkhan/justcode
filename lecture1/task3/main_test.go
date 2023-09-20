package main

import (
	"testing"
)

func TestCompareTwoSlices(t *testing.T) {
	type args struct {
		arr  []int64
		arr1 []int64
	}
	var (
		tests = []struct {
			name string
			args args
			want bool
		}{
			{name: "two slices are equal", args: args{
				[]int64{1, 2, 3},
				[]int64{1, 3, 2},
			}, want: true},
		}
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareTwoSlices(tt.args.arr, tt.args.arr1); got != tt.want {
				t.Errorf("CompareTwoSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
