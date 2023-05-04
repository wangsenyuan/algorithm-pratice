package lcp52

import "sort"

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getNumber(root *TreeNode, ops [][]int) int {
	var nums []int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		nums = append(nums, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	sort.Ints(nums)
	pos := make(map[int]int)
	for i, num := range nums {
		pos[num] = i
	}

	n := len(nums)

	arr := NewNode(0, n-1)
	arr.Update(0, n-1, 0, n-1, -1)

	for _, op := range ops {
		x, y := op[1], op[2]
		x = pos[x]
		y = pos[y]
		var c = op[0]
		if c == 0 {
			c = -1
		}
		arr.Update(x, y, 0, n-1, c)
	}

	var res int
	for i := 0; i < n; i++ {
		x := arr.Get(i, 0, n-1)
		if x == 1 {
			res++
		}
	}

	return res
}

type Node struct {
	left, right *Node
	val         int
}

func NewNode(l int, r int) *Node {
	node := new(Node)
	if l < r {
		mid := (l + r) / 2
		node.left = NewNode(l, mid)
		node.right = NewNode(mid+1, r)
	}
	return node
}

func (node *Node) push(l int, r int) {
	//mid := (l + r) / 2
	if node.left != nil && node.val != 0 {
		node.left.val = node.val
		node.right.val = node.val
		node.val = 0
	}
}

func (node *Node) Update(L int, R int, l int, r int, v int) {
	if R < l || r < L {
		// out of range
		return
	}
	node.push(l, r)
	if L <= l && r <= R {
		node.val = v
		return
	}
	mid := (l + r) / 2
	node.left.Update(L, R, l, mid, v)
	node.right.Update(L, R, mid+1, r, v)
}

func (node *Node) Get(p int, l int, r int) int {
	if l == r {
		return node.val
	}
	node.push(l, r)
	mid := (l + r) / 2
	if p <= mid {
		return node.left.Get(p, l, mid)
	}
	return node.right.Get(p, mid+1, r)
}
