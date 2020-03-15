package p1382

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	hs := make(map[*TreeNode]int)

	rightRotate := func(y *TreeNode) *TreeNode {
		x := y.Left

		t2 := x.Right

		x.Right = y
		y.Left = t2

		hs[y] = max(hs[y.Left], hs[y.Right]) + 1
		hs[x] = max(hs[x.Left], hs[x.Right]) + 1

		return x
	}

	leftRotate := func(x *TreeNode) *TreeNode {
		y := x.Right
		t2 := y.Left

		y.Left = x
		x.Right = t2

		hs[x] = max(hs[x.Left], hs[x.Right]) + 1

		hs[y] = max(hs[y.Left], hs[y.Right]) + 1

		return y
	}

	getBalance := func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return hs[node.Left] - hs[node.Right]
	}

	var Insert func(node *TreeNode, key *TreeNode) *TreeNode

	Insert = func(node *TreeNode, key *TreeNode) *TreeNode {
		if node == nil {
			key.Left = nil
			key.Right = nil
			hs[key] = 1
			return key
		}

		if node.Val > key.Val {
			node.Left = Insert(node.Left, key)
		} else {
			node.Right = Insert(node.Right, key)
		}

		hs[node] = max(hs[node.Left], hs[node.Right]) + 1

		balance := getBalance(node)

		if balance > 1 && key.Val < node.Left.Val {
			return rightRotate(node)
		}

		if balance < -1 && key.Val > node.Right.Val {
			return leftRotate(node)
		}

		if balance > 1 && key.Val > node.Left.Val {
			node.Left = leftRotate(node.Left)
			return rightRotate(node)
		}

		if balance < -1 && key.Val < node.Right.Val {
			node.Right = rightRotate(node.Right)
			return leftRotate(node)
		}

		return node
	}

	var nodes []*TreeNode

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodes = append(nodes, node)
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	var newRoot *TreeNode

	for _, node := range nodes {
		newRoot = Insert(newRoot, node)
	}

	return newRoot
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
