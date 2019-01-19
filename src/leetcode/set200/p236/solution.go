package p236

import "math"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	nodes := make(map[int]*TreeNode)

	open := make(map[int]int)
	close := make(map[int]int)
	PP := make(map[int][]int)
	D := make(map[int]int)
	var dfs func(parent, node *TreeNode, time *int, depth int)
	dfs = func(parent, node *TreeNode, time *int, depth int) {
		*time++
		nodes[node.Val] = node
		open[node.Val] = *time
		PP[node.Val] = make([]int, 20)
		if parent != nil {
			PP[node.Val][0] = parent.Val
		} else {
			PP[node.Val][0] = math.MinInt32
		}

		D[node.Val] = depth

		if node.Left != nil {
			dfs(node, node.Left, time, depth+1)
		}
		if node.Right != nil {
			dfs(node, node.Right, time, depth+1)
		}
		close[node.Val] = *time
	}

	dfs(nil, root, new(int), 0)

	for i := 1; i < 20; i++ {
		for k, _ := range nodes {
			p := PP[k][i-1]
			if p > math.MinInt32 {
				PP[k][i] = PP[p][i-1]
			} else {
				PP[k][i] = math.MinInt32
			}
		}
	}

	isAncester := func(u, v int) bool {
		return open[u] <= open[v] && close[v] <= close[u]
	}

	a := p.Val
	b := q.Val

	if isAncester(a, b) {
		return nodes[a]
	}

	if isAncester(b, a) {
		return nodes[b]
	}

	if D[a] < D[b] {
		a, b = b, a
	}

	//D[a] >= D[b]

	for i := 19; i >= 0; i-- {
		p := PP[a][i]
		if p > math.MinInt32 && !isAncester(p, b) {
			a = p
		}
	}

	return nodes[PP[a][0]]
}

/**
 * Definition for TreeNode.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
