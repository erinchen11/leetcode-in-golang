package findallnumbersdisappearedinanarray

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
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
				nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			},
			want: []int{5, 6},
		},
		{
			name: "example2",
			args: args{
				nums: []int{1, 1},
			},
			want: []int{2},
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := findDisappearedNumbers(v.args.nums); !reflect.DeepEqual(got, v.want){
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}

}
