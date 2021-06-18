package _088_merge_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
	    nums1 []int
	    m int
	    nums2 []int
	    n int
	}


	tests := []struct {
	    name string
	    args args
	    want []int
	}{
	    {
	        name: "Example1",
	        args: args{
	        	nums1: []int{1, 2, 3, 0, 0, 0},
	        	m: 3,
	        	nums2: []int{2, 5, 6},
	        	n: 3,
			},
	        want: []int{1, 2, 2, 3, 5, 6},
	     },
	      {
	        name: "Example2",
	        args: args{
	        	nums1: []int{1},
	        	m: 1,
	        	nums2: []int{},
	        	n: 0,
			},
	        want: []int{1},
	     },
	 }

	 assrt := assert.New(t)
	for _, tt := range tests{
		a, p := tt.want, tt.args

		merge(p.nums1, p.m, p.nums2, p.n)

		assrt.Equal(a, p.nums1)


	}
}
