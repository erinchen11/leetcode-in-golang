package happynumber

import (
	"reflect"
	"testing"
)

func TestIsHappy(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				n: 2,
			},
			want: false,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := isHappy(v.args.n); !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}
}
