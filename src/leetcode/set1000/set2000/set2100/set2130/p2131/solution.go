package p2131

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	var best int
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		best = max(best, arr[i]+arr[j])
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
