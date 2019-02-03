package p987

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, tree *TreeNode, expect [][]int) {
	res := verticalTraversal(tree)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tree := new(TreeNode)
	tree.Val = 3
	tree.Left = new(TreeNode)
	tree.Left.Val = 9
	tree.Right = new(TreeNode)
	tree.Right.Val = 20
	tree.Right.Left = new(TreeNode)
	tree.Right.Left.Val = 15
	tree.Right.Right = new(TreeNode)
	tree.Right.Right.Val = 7
	expect := [][]int{{9}, {3, 15}, {20}, {7}}
	runSample(t, tree, expect)
}

func TestSample2(t *testing.T) {
	tree := new(TreeNode)
	tree.Val = 1
	tree.Left = new(TreeNode)
	tree.Left.Val = 2
	tree.Left.Left = new(TreeNode)
	tree.Left.Left.Val = 4
	tree.Left.Right = new(TreeNode)
	tree.Left.Right.Val = 5
	tree.Right = new(TreeNode)
	tree.Right.Val = 3
	tree.Right.Left = new(TreeNode)
	tree.Right.Left.Val = 6
	tree.Right.Right = new(TreeNode)
	tree.Right.Right.Val = 7
	expect := [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}
	runSample(t, tree, expect)
}
