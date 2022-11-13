package p2471

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

func runSample(t *testing.T, tree string, expect int) {
	res := minimumOperations(ParseAsTree(tree))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tree := "1,4,3,7,6,8,5,null,null,null,null,9,null,10"
	expect := 3
	runSample(t, tree, expect)
}

func TestSample2(t *testing.T) {
	tree := "1,3,2,7,6,5,4"
	expect := 3
	runSample(t, tree, expect)
}

func TestSample3(t *testing.T) {
	tree := "1,2,3,4,5,6"
	expect := 0
	runSample(t, tree, expect)
}
