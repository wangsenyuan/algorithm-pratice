package p1019

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	stack := make([]int, 10010)
	var p int
	ans := make([]int, 10010)
	ii := make([]int, 10010)

	var i int
	node := head
	for node != nil {
		for p > 0 && stack[p-1] < node.Val {
			ans[ii[p-1]] = node.Val
			p--
		}

		stack[p] = node.Val
		ii[p] = i
		p++
		i++
		node = node.Next
	}

	return ans[:i]
}
