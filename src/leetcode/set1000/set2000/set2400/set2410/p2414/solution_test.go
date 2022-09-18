package p2414

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	tree := ParseAsTree(s)

	tree = reverseOddLevels(tree)

	res := SprintTree(tree)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "[2,3,5,8,13,21,34]"
	expect := "[2,5,3,8,13,21,34]"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "[0,1,2,0,0,0,0,1,1,1,1,2,2,2,2]"
	expect := "[0,2,1,0,0,0,0,2,2,2,2,1,1,1,1]"
	runSample(t, s, expect)
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
			if node != nil {
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
