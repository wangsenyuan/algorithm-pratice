package p236

const H = 10

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	indexes := make(map[int]int)
	nodes := make([]*TreeNode, 0, 10)

	var visit func(node *TreeNode)

	visit = func(node *TreeNode) {
		nodes = append(nodes, node)
		i := len(nodes) - 1
		indexes[node.Val] = i
		if node.Left != nil {
			visit(node.Left)
		}
		if node.Right != nil {
			visit(node.Right)
		}
	}
	visit(root)

	n := len(nodes)

	open := make([]int, n)
	close := make([]int, n)
	PP := make([][]int, n)

	for i := 0; i < n; i++ {
		PP[i] = make([]int, H)
	}

	D := make([]int, n)
	var dfs func(parent, node *TreeNode, time *int, depth int)
	dfs = func(parent, node *TreeNode, time *int, depth int) {
		*time++
		i := indexes[node.Val]
		open[i] = *time
		if parent != nil {
			PP[i][0] = indexes[parent.Val]
		} else {
			PP[i][0] = -1
		}

		D[i] = depth

		if node.Left != nil {
			dfs(node, node.Left, time, depth+1)
		}
		if node.Right != nil {
			dfs(node, node.Right, time, depth+1)
		}
		close[i] = *time
	}

	dfs(nil, root, new(int), 0)

	for i := 1; i < H; i++ {
		for k := 0; k < n; k++ {
			p := PP[k][i-1]
			if p >= 0 {
				PP[k][i] = PP[p][i-1]
			} else {
				PP[k][i] = 0
			}
		}
	}

	isAncester := func(u, v int) bool {
		return open[u] <= open[v] && close[v] <= close[u]
	}

	a := indexes[p.Val]
	b := indexes[q.Val]

	if D[a] < D[b] {
		a, b = b, a
	}

	//D[a] >= D[b]

	for i := H - 1; i >= 0; i-- {
		p := PP[a][i]
		if p >= 0 && !isAncester(p, b) {
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
