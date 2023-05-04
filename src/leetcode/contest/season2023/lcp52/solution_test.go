package lcp52

import (
	"strconv"
	"strings"
	"testing"
)

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

func runSample(t *testing.T, s string, ops [][]int, expect int) {
	root := ParseAsTree(s)

	res := getNumber(root, ops)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	root := "[1,null,2,null,3,null,4,null,5]"
	ops := [][]int{
		{1, 2, 4}, {1, 1, 3}, {0, 3, 5},
	}
	expect := 2
	runSample(t, root, ops, expect)
}

func TestSample2(t *testing.T) {
	root := "[4,2,7,1,null,5,null,null,null,null,6]"
	ops := [][]int{
		{0, 2, 2}, {1, 1, 5}, {0, 4, 5}, {1, 5, 7},
	}
	expect := 5
	runSample(t, root, ops, expect)
}
