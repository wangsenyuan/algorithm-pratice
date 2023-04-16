package p2641

import (
	"strconv"
	"strings"
	"testing"
)

func runSample(t *testing.T, root *TreeNode, expect *TreeNode) {
	res := replaceValueInTree(root)

	var dfs func(a *TreeNode, b *TreeNode) bool

	dfs = func(a *TreeNode, b *TreeNode) bool {
		if a == nil && b != nil {
			return false
		}
		if a != nil && b == nil {
			return false
		}
		if a == nil {
			return true
		}
		if a.Val != b.Val {
			return false
		}

		return dfs(a.Left, b.Left) && dfs(a.Right, b.Right)
	}

	if !dfs(expect, res) {
		t.Fatalf("Sample result not correct")
	}
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

func TestSample1(t *testing.T) {
	str := "[5,4,9,1,10,null,7]"
	root := ParseAsTree(str)
	exp := "[0,0,0,7,7,null,11]"
	expect := ParseAsTree(exp)
	runSample(t, root, expect)

}
