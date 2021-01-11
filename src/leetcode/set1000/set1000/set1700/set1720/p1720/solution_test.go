package p1720

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, head *ListNode, k int, expect []int) {
	tmp := swapNodes(head, k)
	res := values(tmp)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}

	k := 2
	expect := []int{1, 4, 3, 2, 5}

	runSample(t, head, k, expect)
}
