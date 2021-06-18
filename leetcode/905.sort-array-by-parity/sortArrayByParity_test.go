package sortarraybyparity

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{
				nums: []int{3, 1, 2, 4},
			},
			want: []int{2, 4, 3, 1},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := sortArrayByParity(v.args.nums); !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}
}
