package lcp60

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

func runSample(t *testing.T, s string, expect int) {
	root := ParseAsTree(s)
	res := getMaxLayerSum(root)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "[6,0,3,null,8]"
	expect := 11
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "[5,6,2,4,null,null,1,3,5]"
	expect := 9
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "[13,97,4,-20,null,null,15,-60,-20,100,200]"
	expect := 300
	runSample(t, s, expect)
}
