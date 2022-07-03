package p2325

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, m int, n int, head *ListNode, expect [][]int) {
	res := spiralMatrix(m, n, head)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func createList(arr []int) *ListNode {
	head := new(ListNode)
	head.Val = arr[0]
	node := head
	for i := 1; i < len(arr); i++ {
		cur := new(ListNode)
		cur.Val = arr[i]
		node.Next = cur
		node = cur
	}
	return head
}

func TestSample1(t *testing.T) {
	head := createList([]int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0})
	m := 3
	n := 5
	expect := [][]int{
		{3, 0, 2, 6, 8}, {5, 0, -1, -1, 1}, {5, 2, 4, 9, 7},
	}
	runSample(t, m, n, head, expect)
}

func TestSample2(t *testing.T) {
	head := createList([]int{0, 1, 2})
	m := 1
	n := 4
	expect := [][]int{
		{0, 1, 2, -1},
	}
	runSample(t, m, n, head, expect)
}
