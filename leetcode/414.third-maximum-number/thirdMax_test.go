package thirdmaximumnumber

import (
	"reflect"
	"testing"
)

func TestThirdMax(t *testing.T) {
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
				nums: []int{3, 2, 1},
			},
			want: 1,
		},
		{
			name: "example2",
			args: args{
				nums: []int{1, 2},
			},
			want: 2,
		},

		{
			name: "example3",
			args: args{
				nums: []int{2, 2, 3, 1},
			},
			want: 1,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := thirdMax(v.args.nums); !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}

}
