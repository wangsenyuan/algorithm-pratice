package p019

import (
	"reflect"
	"testing"
)

func mockData(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	tmp := head
	for i := 1; i < len(nums); i++ {
		node := &ListNode{Val: nums[i]}
		tmp.Next = node
		tmp = node
	}
	return head
}

func runSample(t *testing.T, no int, sz int, head *ListNode, n int, expect []int) {
	newHead := removeNthFromEnd(head, n)
	res := make([]int, sz-1)
	var i int
	for newHead != nil {
		res[i] = newHead.Val
		i++
		newHead = newHead.Next
	}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d with n = %d, should give %v, but got %v", no, n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	head := mockData([]int{1, 2, 3, 4, 5})
	n := 1
	expect := []int{1, 2, 3, 4}
	runSample(t, 1, 5, head, n, expect)
}

func TestSample2(t *testing.T) {
	head := mockData([]int{1, 2, 3, 4, 5})
	n := 2
	expect := []int{1, 2, 3, 5}
	runSample(t, 2, 5, head, n, expect)
}

func TestSample3(t *testing.T) {
	head := mockData([]int{1, 2, 3, 4, 5})
	n := 3
	expect := []int{1, 2, 4, 5}
	runSample(t, 3, 5, head, n, expect)
}

func TestSample4(t *testing.T) {
	head := mockData([]int{1, 2, 3, 4, 5})
	n := 4
	expect := []int{1, 3, 4, 5}
	runSample(t, 4, 5, head, n, expect)
}

func TestSample5(t *testing.T) {
	head := mockData([]int{1, 2, 3, 4, 5})
	n := 5
	expect := []int{2, 3, 4, 5}
	runSample(t, 5, 5, head, n, expect)
}
