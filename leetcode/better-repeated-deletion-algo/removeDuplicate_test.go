package betterrepeateddeletionalgo

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{
				arr: []int{1, 1, 2},
			},
			want: []int{1, 2},
		},
		{
			name: "example2",
			args: args{
				arr: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: []int{0, 1, 2, 3, 4},
		},
	}

	for _, v := range tests {

		t.Run(v.name, func(t *testing.T) {
			if got := v.RemoveDuplicate(args.arr); !reflect.DeepEqual(got, v.want){
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}
}
