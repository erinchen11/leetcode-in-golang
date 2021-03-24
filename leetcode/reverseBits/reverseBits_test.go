package leetcode

import "testing"

/*
Reverse bits of a given 32 bits unsigned integer

思路：先將input轉成無號整數，再轉成32位進制
*/

func Test_reverseBits(t *testing.T) {
	tests := []struct {
		name string
		want uint32
		n    uint32
	}{
		{
			name: "example1",
			want: 964176192,
			n:    43261596,
		},
		{
			name: "example2",
			want: 3221225471,
			n:    4294967293,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.n); got != tt.want {
				t.Errorf("reverseBits = %v want %v\n", got, tt.want)
			}
		})
	}
}
