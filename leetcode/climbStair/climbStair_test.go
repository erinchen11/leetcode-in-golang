package leetcode

import "testing"

/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps.
In how many distinct ways can you climb to the top?
*/

func Test_climbStair(t *testing.T) {
	tests := []struct {
		name string
		want int
		n    int
	}{
		{
			name: "example1",
			n:    4,
			want: 5,
		},

		{
			name: "example2",
			n:    3,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStair(tt.n); got != tt.want {
				t.Errorf("climbStair = %v want %v\n", got, tt.want)
			}
		})
	}

}
