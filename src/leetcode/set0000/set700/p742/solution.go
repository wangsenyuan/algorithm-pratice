package p742

import . "leetcode/util"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findClosestLeaf(root *TreeNode, k int) int {
	grid := make(map[int]map[int]bool)
	leaf := make(map[int]bool)
	var visit func(parent, node *TreeNode)

	visit = func(parent, node *TreeNode) {
		grid[node.Val] = make(map[int]bool)
		if parent != nil {
			grid[node.Val][parent.Val] = true
		}

		if node.Left == nil && node.Right == nil {
			leaf[node.Val] = true
			return
		}

		if node.Left != nil {
			grid[node.Val][node.Left.Val] = true
			visit(node, node.Left)
		}

		if node.Right != nil {
			grid[node.Val][node.Right.Val] = true
			visit(node, node.Right)
		}
	}
	visit(nil, root)

	n := len(grid)

	if n == 1 {
		return k
	}

	visited := make(map[int]bool)
	visited[k] = true

	bfs := func(k int) int {
		que := make([]int, n)
		head, tail := 0, 0
		que[tail] = k
		tail++
		dist := 0
		for head < tail {
			tmp := tail
			for head < tmp {
				cur := que[head]
				visited[cur] = true
				head++
				if leaf[cur] {
					// a leaf
					return cur
				}
				for w := range grid[cur] {
					if !visited[w] {
						que[tail] = w
						tail++
					}
				}
			}
			dist++
		}
		return -1
	}
	return bfs(k)
}
