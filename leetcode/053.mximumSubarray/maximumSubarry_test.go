package leetcode

import (
	"fmt"
	"testing"
)

/*
Given an integer array nums, find the contiguous subarray
(containing at least one number)
which has the largest sum and return its sum
*/

func Test_maximumSubarray(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "example1",
			args: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "example2",
			args: []int{1},
			want: 1,
		},
		{
			name: "example3",
			args: []int{-10000},
			want: -10000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSubarray(tt.args); got != tt.want {
				fmt.Errorf("maximumSubarray = %v  want %v\n", got, tt.want)
			}
		})
	}

}
