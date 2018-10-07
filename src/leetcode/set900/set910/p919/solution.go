package p919

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

type CBTInserter struct {
	root  *TreeNode
	nodes []*TreeNode
	pp    int
	next  []*TreeNode
	pos   int
}

func Constructor(root *TreeNode) CBTInserter {
	rootNode := new(TreeNode)
	rootNode.Val = root.Val
	nodes := []*TreeNode{rootNode}
	next := make([]*TreeNode, 2)
	sol := CBTInserter{rootNode, nodes, 0, next, 0}

	level := make([]*TreeNode, 0, 10)
	level = append(level, root)
	var cur int
	for cur < len(level) {
		node := level[cur]
		cur++
		if node != root {
			sol.Insert(node.Val)
		}
		if node.Left != nil {
			level = append(level, node.Left)
		}
		if node.Right != nil {
			level = append(level, node.Right)
		}
	}

	return sol
}

func (this *CBTInserter) Insert(v int) int {
	parent := this.nodes[this.pp]

	node := new(TreeNode)
	node.Val = v

	this.next[this.pos] = node
	this.pos++

	if this.pos&1 == 0 {
		parent.Right = node
		//advance parent
		this.pp++
		if this.pp == len(this.nodes) {
			//full
			this.nodes = this.next
			this.next = make([]*TreeNode, 2*len(this.nodes))
			this.pp = 0
			this.pos = 0
		}
	} else {
		parent.Left = node
	}

	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
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
