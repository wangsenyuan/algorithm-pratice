package p889_

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) != len(post) {
		return nil
	}
	if len(pre) == 0 {
		return nil
	}
	n := len(pre)
	root := new(TreeNode)
	root.Val = pre[0]
	if n == 1 {
		return root
	}
	a := make(map[int]bool)
	b := make(map[int]bool)

	for i := 1; i < n; i++ {
		a[pre[i]] = true
		b[post[i-1]] = true
		if diffCount(a, b) == 0 {
			root.Left = constructFromPrePost(pre[1:i+1], post[:i])
			root.Right = constructFromPrePost(pre[i+1:], post[i:n-1])
			break
		}
	}
	return root
}

func diffCount(a, b map[int]bool) int {
	if len(a) > len(b) {
		return len(a) - len(b)
	}
	if len(a) < len(b) {
		return len(b) - len(a)
	}
	var cnt int
	for k := range a {
		if !b[k] {
			cnt++
		}
	}
	for k := range b {
		if !a[k] {
			cnt++
		}
	}
	return cnt
}

func preOrderTraversal(root *TreeNode) []int {
	var travel func(root *TreeNode)

	var result []int
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		travel(root.Left)
		travel(root.Right)
	}

	travel(root)

	return result
}

func postOrderTraversal(root *TreeNode) []int {
	var travel func(root *TreeNode)

	var result []int
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		travel(root.Right)
		result = append(result, root.Val)
	}

	travel(root)

	return result
}
