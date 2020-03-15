package p1382

import (
	"testing"
)

func TestSample1(t *testing.T) {
	root := new(TreeNode)
	root.Val = 1
	root.Right = new(TreeNode)
	root.Right.Val = 2
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 3
	root.Right.Right.Right = new(TreeNode)
	root.Right.Right.Right.Val = 4
	res := balanceBST(root)

	t.Logf("%d", res.Val)
}
