package p742

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) String() string {
	if tn == nil {
		return "null"
	}
	return fmt.Sprintf("%d,%v,%v", tn.Val, tn.Left, tn.Right)
}

func ParseAsTree(str string) *TreeNode {
	str = strings.Replace(str, "[", "", 1)
	str = strings.Replace(str, "]", "", 1)

	var currLevel []*TreeNode
	var nextLevel []*TreeNode
	nodes := strings.Split(str, ",")

	root := parseNode(nodes[0])
	currLevel = append(currLevel, root)

	pt := 0
	ct := 0
	for i := 1; i < len(nodes); i++ {
		parent := currLevel[pt]
		node := parseNode(nodes[i])
		if node != nil {
			nextLevel = append(nextLevel, node)
		}
		if ct == 0 {
			parent.Left = node
			ct++
		} else {
			parent.Right = node
			pt++
			ct = 0
		}
		if pt == len(currLevel) {
			currLevel = nextLevel
			nextLevel = make([]*TreeNode, 0)
			pt = 0
		}
	}
	return root
}

func parseNode(str string) *TreeNode {
	if str == "null" {
		return nil
	}
	return &TreeNode{parseNum(str), nil, nil}
}
func parseNum(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

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
