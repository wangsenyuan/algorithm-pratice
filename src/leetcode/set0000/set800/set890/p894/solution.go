package p894

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(N int) []*TreeNode {
	if N == 0 {
		return nil
	}

	if N&1 == 0 {
		return nil
	}
	if N == 1 {
		root := new(TreeNode)
		return []*TreeNode{root}
	}
	res := make([]*TreeNode, 0, 10)

	for i := 1; i < N; i += 2 {
		// rooted at i
		ls := allPossibleFBT(i)
		if len(ls) == 0 {
			continue
		}
		rs := allPossibleFBT(N - i - 1)
		if len(rs) == 0 {
			continue
		}
		for _, l := range ls {
			for _, r := range rs {
				root := new(TreeNode)
				root.Left = l
				root.Right = r
				res = append(res, root)
			}
		}
	}

	return res
}
