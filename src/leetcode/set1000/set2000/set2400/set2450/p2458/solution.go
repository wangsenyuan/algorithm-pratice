package p2458

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	n := size(root)

	in := make([]int, n)
	out := make([]int, n)
	tour := make([]int, 0, 2*n)
	D := make([]int, n)

	var dfs func(node *TreeNode, depth int)

	dfs = func(node *TreeNode, depth int) {
		u := node.Val - 1
		D[u] = depth
		in[u] = len(tour)
		tour = append(tour, u)

		if node.Left != nil {
			dfs(node.Left, depth+1)
		}

		if node.Right != nil {
			dfs(node.Right, depth+1)
		}
		out[u] = len(tour)
		tour = append(tour, u)
	}

	dfs(root, 0)

	m := len(tour)

	arr := make([]int, 2*m)

	for i := m; i < 2*m; i++ {
		arr[i] = D[tour[i-m]]
	}

	for i := m - 1; i > 0; i-- {
		arr[i] = max(arr[i*2], arr[i*2+1])
	}

	get := func(l int, r int) int {
		l += m
		r += m
		var res int
		for l < r {
			if l&1 == 1 {
				res = max(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	ans := make([]int, len(queries))

	for i, u := range queries {
		u--
		if in[u] > 0 {
			ans[i] = get(0, in[u])
		}
		if out[u]+1 < m {
			ans[i] = max(ans[i], get(out[u]+1, m))
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func size(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + size(node.Left) + size(node.Right)
}
