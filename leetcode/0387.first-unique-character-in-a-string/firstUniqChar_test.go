package firstuniquecharacterinastring

import (
	"reflect"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
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
			args: args{
				s: "leetcode",
			},
			want: 0,
		},
		{
			name: "example2",
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
		{
			name: "example3",
			args: args{
				s: "aabb",
			},
			want: -1,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := firstUniqChar(v.args.s); !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v, want %v ", got, v.want)
			}
		})
	}
}
