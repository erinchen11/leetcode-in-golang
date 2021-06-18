package singlenumber

import (
	"reflect"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				nums: []int{2, 2, 1},
			},
			want: 1,
		},
		{
			name: "example2",
			args: args{
				nums: []int{4, 1, 2, 1, 2},
			},
			want: 4,
		},
		{
			name: "example2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := singleNumber(v.args.nums); !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}
}
