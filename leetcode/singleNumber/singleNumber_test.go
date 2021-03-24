package leetcode

import "testing"

/*
Given a non-empty array of integers nums,
every element appears twice except for one.
Find that single one.

Follow up: Could you implement a solution with a linear runtime
complexity and without using extra memory?
*/

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example1",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "example2",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "example3",
			nums: []int{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.nums); got != tt.want {
				t.Errorf("singleNumber = %v want %v\n", got, tt.want)
			}
		})
	}
}
