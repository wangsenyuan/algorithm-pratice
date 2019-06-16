package p328

import "testing"

func runSample(t *testing.T, no int, head *ListNode, expect *ListNode) {
	res := oddEvenList(head)
	for expect != nil && res != nil {
		if expect.Val != res.Val {
			t.Errorf("sample %d, expect: %d, res: %d", no, expect.Val, res.Val)
		}
		expect = expect.Next
		res = res.Next
	}
	if expect != nil || res != nil {
		t.Errorf("sample %d, different length result", no)
	}
}

func mockData(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	tmp := head
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{Val: nums[i]}
		tmp = tmp.Next
	}
	return head
}

func TestSample1(t *testing.T) {
	head := mockData([]int{1, 2, 3, 4, 5, 6})
	expect := mockData([]int{1, 3, 5, 2, 4, 6})
	runSample(t, 1, head, expect)
}
