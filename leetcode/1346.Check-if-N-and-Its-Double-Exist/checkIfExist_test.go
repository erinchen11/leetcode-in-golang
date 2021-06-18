package checkifnanditsdoubleexist

import (
	"reflect"
	"testing"
)

func TestCheckIfExist(t *testing.T) {
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
				nums: []int{10, 2, 5, 3},
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				nums: []int{3, 1, 7, 11},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckIfExist(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
