package leetcode

import (
	"reflect"
	"testing"
)

/*

Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice.

You can return the answer in any order.

input: []int, int
output: []int
*/
// write table driven test

// type question struct {
// 	input
// 	output
// }

// type input struct {
// 	nums   []int
// 	target int
// }

// type output struct {
// 	output []int
// }

// func TestTwoSum(t *testing.T) {
// 	nums := []int{2, 7, 14, 18}
// 	target := 9
// 	res := TwoSum(nums, target)
// 	fmt.Println(res)

// }

// func TestTwoSum(t *testing.T) {
// 	nums := []int{2, 7, 14, 18}
// 	target := 9
// 	res := TwoSum(nums, target)
// 	fmt.Println(res)
// }

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "Example2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 5,
			},
			want: []int{0, 1},
		},
		{
			name: "Example3",
			args: args{
				nums:   []int{0, 8, 7, 3, 3, 4, 2},
				target: 11,
			},
			want: []int{1, 3},
		},
		{
			name: "Example4",
			args: args{
				nums:   []int{0, 1},
				target: 1,
			},
			want: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v , want %v", got, tt.want)
			}
		})
	}

}
