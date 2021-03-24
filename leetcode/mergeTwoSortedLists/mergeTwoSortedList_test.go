package leetcode

import "testing"

/*
Merge two sorted linked lists and return it as a sorted list.
The list should be made by splicing together the nodes of the
first two lists.

*/

func Test_mergeTwoList(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{
				a: []int{1, 2, 4},
				b: []int{1, 3, 4},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "example2",
			args: args{
				a: []int{},
				b: []int{0},
			},
			want: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("mergeTwoList = %v want %v", got, tt.want)
			}
		})
	}
}
