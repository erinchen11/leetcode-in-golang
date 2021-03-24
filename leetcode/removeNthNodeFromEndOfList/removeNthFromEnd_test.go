package leetcode

import "testing"

/*
Given the head of a linked list, remove the nth node from the
end of the list and return its head.

Follow up: Could you do this in one pass?
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }


*/

func Test_removeNthFromEnd(t *testing.T) {

	type ListNode struct {
		Val int
		Next *ListNode
	}
	type args struct {
		head ListNode
		n int
	}
	tests := []struct{
		name string
		args args
		want *ListNode
	}{
		{
			name: "example1",
			args: args{
				head : {
					Val : 
				},
				n : 1,
			},
			want: [],
		},
	}

}
