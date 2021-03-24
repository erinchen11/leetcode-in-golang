package leetcode

import "testing"

/*
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters
in s can be replaced to get t.

All occurrences of a character must be replaced with
another character while preserving the order of characters.
No two characters may map to the same character,
but a character may map to itself.
*/

func Test_isIsomorphics(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				s: "egg",
				t: "add",
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				s: "foo",
				t: "bar",
			},
			want: false,
		},
		{
			name: "example3",
			args: args{
				s: "paper",
				t: "title",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphics(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphics = %v, want %v", got, tt.want)
			}
		})
	}
}
