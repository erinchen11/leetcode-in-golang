package movezeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T){
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
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "example2",
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := MoveZeroes(v.args.nums); !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}

}