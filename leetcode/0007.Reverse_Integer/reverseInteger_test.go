package leetcode

import (
	"reflect"
	"testing"
)

func Test_Reverse(t *testing.T) {
	type args struct {
		input int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				input: 123,
			},
			want: 321,
		},
		{
			name: "Example2",
			args: args{
				input: 120,
			},
			want: 21,
		},
		{
			name: "Example3",
			args: args{
				input: -123,
			},
			want: -321,
		},
		{
			name: "Example4",
			args: args{
				input: 0,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse function is %v want is %v", got, tt.want)
			}
		})
	}
}
