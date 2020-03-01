package p1367

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	if contains(head, root) {
		return true
	}

	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func contains(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}

	if root == nil || root.Val != head.Val {
		return false
	}

	return contains(head.Next, root.Left) || contains(head.Next, root.Right)
}
