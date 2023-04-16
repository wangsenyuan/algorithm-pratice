package p2641

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	var sum []int

	var dfs func(node *TreeNode, d int)
	dfs = func(node *TreeNode, d int) {
		if node == nil {
			return
		}
		if len(sum) == d {
			sum = append(sum, node.Val)
		} else {
			sum[d] += node.Val
		}
		dfs(node.Left, d+1)
		dfs(node.Right, d+1)
	}

	dfs(root, 0)

	var dfs2 func(node *TreeNode, d int) *TreeNode

	dfs2 = func(node *TreeNode, d int) *TreeNode {
		if node == nil {
			return nil
		}
		res := new(TreeNode)
		res.Val = sum[d] - node.Val
		res.Left = dfs2(node.Left, d+1)
		if node.Right != nil && res.Left != nil {
			res.Left.Val -= node.Right.Val
		}
		res.Right = dfs2(node.Right, d+1)
		if node.Left != nil && res.Right != nil {
			res.Right.Val -= node.Left.Val
		}
		return res
	}

	return dfs2(root, 0)
}
