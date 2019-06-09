package p5084

import (
	"bytes"
	"math"
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

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}

	var loop func(node *TreeNode, sum int) (int, *TreeNode)

	loop = func(node *TreeNode, sum int) (int, *TreeNode) {
		x := new(TreeNode)
		x.Val = node.Val

		if node.Left == nil && node.Right == nil {
			if node.Val+sum < limit {
				return x.Val, nil
			}

			return x.Val, x
		}
		cur := math.MinInt32
		if node.Left != nil {
			a, l := loop(node.Left, sum+node.Val)
			x.Left = l
			cur = a + node.Val
		}

		if node.Right != nil {
			b, r := loop(node.Right, sum+node.Val)
			x.Right = r
			cur = max(cur, b+node.Val)
		}

		if cur+sum >= limit {
			return cur, x
		}

		return cur, nil
	}

	_, r := loop(root, 0)

	return r
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
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

func SprintTree(root *TreeNode) string {
	if root == nil {
		return "[null]"
	}
	var buf bytes.Buffer
	buf.WriteByte('[')
	curLevel := make([]*TreeNode, 0)
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		nextLevel := make([]*TreeNode, 0)
		for _, node := range curLevel {
			if node == nil {
				buf.WriteString("null,")
			} else {
				buf.WriteString(strconv.Itoa(node.Val))
				buf.WriteString(",")
				nextLevel = append(nextLevel, node.Left)
				nextLevel = append(nextLevel, node.Right)
			}
		}
		curLevel = nextLevel
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}

	buf.WriteByte(']')

	return buf.String()
}
