package p2458

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func runSample(t *testing.T, str string, queries []int, expect []int) {
	tree := ParseAsTree(str)

	res := treeQueries(tree, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
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
	str := "[1,3,4,2,null,6,5,null,null,null,null,null,7]"
	queries := []int{4}
	expect := []int{2}
	runSample(t, str, queries, expect)
}

func TestSample2(t *testing.T) {
	str := "[5,8,9,2,1,3,7,4,6]"
	queries := []int{3, 2, 4, 8}
	expect := []int{3, 2, 3, 2}
	runSample(t, str, queries, expect)
}
