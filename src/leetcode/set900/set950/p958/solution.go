package p958

import (
	"bytes"
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

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := make([]*TreeNode, 0, 128)
	root.Val = 1
	res = append(res, root)
	var i int
	for i < len(res) {
		cur := res[i]
		i++
		if cur.Left != nil {
			cur.Left.Val = 2 * i
			res = append(res, cur.Left)
		}
		if cur.Right != nil {
			cur.Right.Val = 2*i + 1
			res = append(res, cur.Right)
		}
	}
	return res[i-1].Val == len(res)
}

func isCompleteTree1(root *TreeNode) bool {
	var visit func(root *TreeNode, depth int) (int, int)

	visit = func(root *TreeNode, depth int) (int, int) {
		if root == nil {
			return depth, depth
		}
		a, b := visit(root.Left, depth+1)
		c, d := visit(root.Right, depth+1)
		if a < 0 || c < 0 {
			return -1, -1
		}
		if a < b || c < d {
			return -1, -1
		}
		if a > b+1 || c > d+1 {
			return -1, -1
		}

		if a == b {
			if a == c {
				return a, d
			}
			if c == d && a == c+1 {
				return a, d
			}
			return -1, -1
		}
		// a > b
		if b == c {
			if c == d {
				return a, d
			}
		}
		return -1, -1
	}

	a, b := visit(root, 0)
	if a < 0 {
		return false
	}
	return a == b || a == b+1
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

	return buf.String()
}
