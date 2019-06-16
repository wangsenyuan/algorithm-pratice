package p987

import "sort"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	xi := make(map[*TreeNode]int)
	yi := make(map[*TreeNode]int)

	var dfs func(node *TreeNode, x int, y int)

	dfs = func(node *TreeNode, x int, y int) {
		if node == nil {
			return
		}
		xi[node] = x
		yi[node] = y
		dfs(node.Left, x-1, y-1)
		dfs(node.Right, x+1, y-1)
	}

	dfs(root, 0, 0)

	var a int
	var b int
	for _, x := range xi {
		if x < a {
			a = x
		}
		if x > b {
			b = x
		}
	}

	res := make([][]Pair, b-a+1)

	for i := 0; i < len(res); i++ {
		res[i] = make([]Pair, 0, 10)
	}

	for node, x := range xi {
		i := x - a
		res[i] = append(res[i], Pair{yi[node], node.Val})
	}
	ans := make([][]int, len(res))

	for i := 0; i < len(res); i++ {
		sort.Sort(Pairs(res[i]))
		ans[i] = make([]int, len(res[i]))
		for j := 0; j < len(res[i]); j++ {
			ans[i][j] = res[i][j].second
		}
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	a, b := this[i], this[j]
	return a.first > b.first || (a.first == b.first && a.second < b.second)
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
