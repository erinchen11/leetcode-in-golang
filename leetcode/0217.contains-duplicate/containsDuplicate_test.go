package containsduplicate

import (
	"reflect"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "example2",
			args: args{
				nums: []int{1, 2, 1, 3, 3, 4},
			},
			want: true,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := containsDuplicate(v.args.nums); !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}
}
