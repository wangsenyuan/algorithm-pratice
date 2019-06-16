package p965

import "testing"

func runSample(t *testing.T, root *TreeNode, expect bool) {
	res := isUnivalTree(root)

	if res != expect {
		t.Errorf("Sample should get %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	root := new(TreeNode)
	root.Val = 2
	root.Left = new(TreeNode)
	root.Left.Val = 3
	runSample(t, root, false)
}

func TestSample2(t *testing.T) {
	root := new(TreeNode)
	root.Val = 2
	root.Left = new(TreeNode)
	root.Left.Val = 2
	runSample(t, root, true)
}
