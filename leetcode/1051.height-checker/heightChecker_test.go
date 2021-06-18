package heightchecker

import (
	"reflect"
	"testing"
)

func TestHeightChecker(t *testing.T) {
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
				nums: []int{5,1,2,3,4},
			},
			want: 5,
		},
		{
			name: "example2",
			args: args{
				nums: []int{1,1,4,2,1,3},
			},
			want:3 ,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := heightChecker(v.args.nums); !reflect.DeepEqual(got, v.want){
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}
}
