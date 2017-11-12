package p725

import . "leetcode/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
	sz := size(root)

	a, b := sz/k, sz%k

	res := make([]*ListNode, k)

	head := root
	for i := 0; i < k && head != nil; i++ {
		res[i] = head
		tmp := head
		for j := 0; j < a; j++ {
			tmp = head
			head = head.Next
		}
		if b > 0 {
			tmp = head
			head = head.Next
			b--
		}
		tmp.Next = nil
	}
	return res
}

func size(root *ListNode) int {
	res := 0
	tmp := root
	for tmp != nil {
		res++
		tmp = tmp.Next
	}
	return res
}
