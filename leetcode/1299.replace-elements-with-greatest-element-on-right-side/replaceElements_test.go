package replaceelementswithgreatestelementonrightside

import (
	"reflect"
	"testing"
)

func TestReplaceElements(t *testing.T) {
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
				nums: []int{17, 18, 5, 4, 6, 1},
			},
			want: []int{18, 6, 6, 6, 1, -1},
		},
		{
			name: "example2",
			args: args{
				nums: []int{400},
			},
			want: []int{-1},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := ReplaceElements(v.args.nums); !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}
}
