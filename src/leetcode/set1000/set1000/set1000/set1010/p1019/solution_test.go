package p1019

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, arr []int, expect []int) {
	head := new(ListNode)
	head.Val = arr[0]
	node := head
	for i := 1; i < len(arr); i++ {
		node.Next = new(ListNode)
		node.Next.Val = arr[i]
		node = node.Next
	}
	res := nextLargerNodes(head)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{2, 1, 5}
	expect := []int{5, 5, 0}

	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{2, 7, 4, 3, 5}
	expect := []int{7, 0, 5, 5, 0}

	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 7, 5, 1, 9, 2, 5, 1}
	expect := []int{7, 9, 9, 9, 0, 5, 0, 0}

	runSample(t, arr, expect)
}
