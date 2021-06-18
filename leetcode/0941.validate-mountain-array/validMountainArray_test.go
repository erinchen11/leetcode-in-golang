package validatemountainarray

import (
	"reflect"
	"testing"
)

func TestValidMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				arr: []int{2, 1},
			},
			want: false,
		},
		{
			name: "example2",
			args: args{
				arr: []int{3, 5, 5},
			},
			want: false,
		},
		{
			name: "example3",
			args: args{
				arr: []int{0, 3, 2, 1},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidMountainArray(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
