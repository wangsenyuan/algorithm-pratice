package a

import "testing"

func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	node := new(ListNode)
	node.Val = arr[0]
	node.Next = createList(arr[1:])
	return node
}

func runSample(t *testing.T, arr []int, expect bool) {
	head := createList(arr)
	res := isPalindrome(head)
	if res != expect {
		t.Errorf("Sample %v expect %t, but got %t", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expect := false
	runSample(t, arr, expect)
}
