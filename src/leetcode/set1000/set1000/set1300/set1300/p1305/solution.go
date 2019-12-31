package p1305

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	nums1 := getNums(root1)
	nums2 := getNums(root2)
	res := make([]int, 0, len(nums1)+len(nums2))
	var i, j int

	for i < len(nums1) || j < len(nums2) {
		if i == len(nums1) || (j < len(nums2) && nums1[i] >= nums2[j]) {
			res = append(res, nums2[j])
			j++
		} else {
			res = append(res, nums1[i])
			i++
		}
	}
	return res
}

func getNums(root *TreeNode) []int {
	var res []int

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return res
}
