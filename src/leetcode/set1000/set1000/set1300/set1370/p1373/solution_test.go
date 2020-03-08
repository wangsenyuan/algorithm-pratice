package p1373

import (
	"strconv"
	"strings"
	"testing"
)

func TestSample1(t *testing.T) {
	s := "1,4,3,2,4,2,5,null,null,null,null,null,null,4,6"
	arr := strings.Split(s, ",")
	tree := ParseAsTree(arr)
	expect := 20
	res := maxSumBST(tree)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample2(t *testing.T) {
	s := "4,3,null,1,2"
	arr := strings.Split(s, ",")
	tree := ParseAsTree(arr)
	expect := 2
	res := maxSumBST(tree)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample3(t *testing.T) {
	s := "-4,-2,-5"
	arr := strings.Split(s, ",")
	tree := ParseAsTree(arr)
	expect := 0
	res := maxSumBST(tree)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func ParseAsTree(arr []string) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	trees := make([]*TreeNode, len(arr))
	var p int
	var k int
	for i := 0; i < len(arr); i++ {
		var cur *TreeNode
		if arr[i] != "null" {
			cur = new(TreeNode)
			cur.Val, _ = strconv.Atoi(arr[i])
		}

		trees[i] = cur
		if i > p {
			k++
			if k == 1 {
				trees[p].Left = cur
			} else {
				trees[p].Right = cur
				p++
				k = 0
			}
		}
	}
	return trees[0]
}
