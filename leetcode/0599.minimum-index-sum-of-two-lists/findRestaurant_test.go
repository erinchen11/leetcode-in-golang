package minimumindexsumoftwolists

import (
	"reflect"
	"testing"
)

func TestFindRestaurant(t *testing.T) {
	type args struct {
		list1 []string
		list2 []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example1",
			args: args{
				list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				list2: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			},
			want: []string{"Shogun"},
		},
		{
			name: "example2",
			args: args{
				list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				list2: []string{"KFC", "Shogun", "Burger King"},
			},
			want: []string{"Shogun"},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := findRestaurant(v.args.list1, v.args.list2); !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}
}
