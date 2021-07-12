package p1932

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const N = 500010

func canMerge(trees []*TreeNode) *TreeNode {
	A := make([]int, N)
	B := make([]int, N)
	for i := 1; i < N; i++ {
		A[i] = N
		B[i] = 0
	}
	X := make([]*TreeNode, N)
	in := make([]bool, N)
	for _, t := range trees {
		X[t.Val] = t
		if t.Left != nil {
			in[t.Left.Val] = true
		}
		if t.Right != nil {
			in[t.Right.Val] = true
		}
	}

	var root *TreeNode

	for i := 1; i < N; i++ {
		if X[i] != nil {
			if !in[i] {
				if root != nil {
					// multiple roots
					return nil
				}
				root = X[i]
			}
		}
	}

	if root == nil {
		return nil
	}

	A[root.Val] = 0
	B[root.Val] = N

	var dfs func(node *TreeNode) *TreeNode

	var cnt int

	dfs = func(node *TreeNode) *TreeNode {
		cnt++

		res := new(TreeNode)
		res.Val = node.Val

		if node.Left != nil {
			if node.Left.Val <= A[node.Val] {
				return nil
			}
			// node.Left.Val > A[node.val]
			A[node.Left.Val] = A[node.Val]
			B[node.Left.Val] = node.Val
			if X[node.Left.Val] != nil {
				tmp := dfs(X[node.Left.Val])
				if tmp == nil {
					return nil
				}
				res.Left = tmp
			} else {
				res.Left = node.Left
			}
		}
		if node.Right != nil {
			if node.Right.Val >= B[node.Val] {
				return nil
			}
			A[node.Right.Val] = node.Val
			B[node.Right.Val] = B[node.Val]
			if X[node.Right.Val] != nil {
				tmp := dfs(X[node.Right.Val])
				if tmp == nil {
					return nil
				}
				res.Right = tmp
			} else {
				res.Right = node.Right
			}
		}
		return res
	}
	res := dfs(root)

	if cnt != len(trees) {
		return nil
	}

	return res
}
