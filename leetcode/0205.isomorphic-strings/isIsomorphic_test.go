package isomorphicstrings

import (
	"reflect"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {

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

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := isIsomorphic(v.args.s, v.args.t); !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}
}
