package maximumsubarray

import (
	"reflect"
	"testing"
)

func Test_maxSubArray(t *testing.T)  {
	

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
				nums: []int{-2,1,-3,4,-1,2,1,-5,4},
			},
			want: 6,
		},
		{
			name: "example2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "example2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); !reflect.DeepEqual(got, tt.want){
				t.Errorf("Got %v want %v", got, tt.want)
			}
		})
	}
}