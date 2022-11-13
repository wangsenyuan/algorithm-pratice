package p2471

import "sort"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumOperations(root *TreeNode) int {

	var layers [][]int

	add := func(d int, v int) {
		if len(layers) == d {
			layers = append(layers, []int{v})
		} else {
			layers[d] = append(layers[d], v)
		}
	}

	var n int

	var dfs func(node *TreeNode, d int)

	dfs = func(node *TreeNode, d int) {
		n++
		add(d, node.Val)
		if node.Left != nil {
			dfs(node.Left, d+1)
		}
		if node.Right != nil {
			dfs(node.Right, d+1)
		}
	}

	dfs(root, 0)

	var res int

	buf := make([]int, n)
	vis := make([]bool, n)

	for _, lay := range layers {
		res += swap(lay, buf[:len(lay)], vis)
	}

	return res
}

func swap(arr []int, buf []int, vis []bool) int {
	copy(buf, arr)
	sort.Ints(buf)

	pos := make(map[int]int)

	for i, num := range arr {
		pos[num] = i
		vis[i] = false
	}

	var res int
	for i := 0; i < len(buf); i++ {
		if vis[i] {
			continue
		}
		j := i
		for !vis[j] {
			vis[j] = true
			num := buf[j]
			j = pos[num]
			res++
		}
		res--
	}

	return res
}
