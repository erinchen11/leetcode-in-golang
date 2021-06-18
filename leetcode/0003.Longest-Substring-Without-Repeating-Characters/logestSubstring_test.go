package leetcode

import (
	"reflect"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	// parameters of input
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "example2",
			args: args{s: "bbbbb"},
			want: 1,
		},
		{
			name: "example3",
			args: args{s: "pwwkew"},
			want: 3,
		},
		{
			name: "example4",
			args: args{s: ""},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lengthOfLongestSubstring() = %v, want = %v", got, tt.want)
			}
		})
	}
}
