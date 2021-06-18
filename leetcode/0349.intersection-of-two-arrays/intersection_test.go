package intersectionoftwoarrays

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			want: []int{2},
		},
		{
			name: "example2",
			args: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			want: []int{9, 4},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := intersection(v.args.nums1, v.args.nums2); !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}
}
