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
	//arr.Update(0, n-1, 0, n-1, 0)

	for _, op := range ops {
		x, y := op[1], op[2]
		x = pos[x]
		y = pos[y]
		var c = op[0]

		arr.Update(x, y, 0, n-1, c)
	}

	return arr.cnt
}

type Node struct {
	left, right *Node
	val         int
	cnt         int
}

func NewNode(l int, r int) *Node {
	node := new(Node)
	node.val = -1
	if l < r {
		mid := (l + r) / 2
		node.left = NewNode(l, mid)
		node.right = NewNode(mid+1, r)
	}
	return node
}

func (node *Node) pushVal(l int, r int, v int) {
	node.val = v
	node.cnt = v * (r - l + 1)
}

func (node *Node) push(l int, r int) {
	mid := (l + r) / 2
	if node.left != nil && node.val != -1 {
		if node.left.val != node.val {
			node.left.pushVal(l, mid, node.val)
		}
		if node.right.val != node.val {
			node.right.pushVal(mid+1, r, node.val)
		}
		node.val = -1
	}
}

func (node *Node) Update(L int, R int, l int, r int, v int) {
	if R < l || r < L {
		// out of range
		return
	}
	node.push(l, r)
	if L <= l && r <= R {
		node.pushVal(l, r, v)
		return
	}
	mid := (l + r) / 2
	node.left.Update(L, R, l, mid, v)
	node.right.Update(L, R, mid+1, r, v)
	node.cnt = node.left.cnt + node.right.cnt
}
