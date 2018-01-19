package p103

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	tree := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	result := zigzagLevelOrder(tree)

	expect := [][]int{
		{3},
		{20, 9},
		{15, 7},
	}

	i := 0
	for i < len(result) && i < len(expect) {
		if !reflect.DeepEqual(result[i], expect[i]) {
			t.Errorf("sample has wrong answer: %v vs %v", result[i], expect[i])
		}
		i++
	}
	if i < len(result) || i < len(expect) {
		t.Errorf("sample has wrong answer: %v vs %v", result, expect)
	}
}
