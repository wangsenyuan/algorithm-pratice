package p1161

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	var nums []int

	var visit func(node *TreeNode, level int)

	visit = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(nums) > level {
			nums[level] += node.Val
		} else {
			nums = append(nums, node.Val)
		}

		visit(node.Left, level+1)
		visit(node.Right, level+1)
	}

	visit(root, 0)

	var ans int

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[ans] {
			ans = i
		}
	}

	return ans + 1
}
