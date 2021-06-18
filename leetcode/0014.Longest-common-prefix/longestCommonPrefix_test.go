package leetcode

import (
	"reflect"
	"testing"
)

//Write a function to find the longest common prefix string
//amongst an array of strings.

//If there is no common prefix, return an empty string "".

// input []string
// output string

func Test_LongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "Example2",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
