package p968

import "testing"

func runSample(t *testing.T, root *TreeNode, expect int) {
	res := minCameraCover(root)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	root := &TreeNode{0, &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, nil}
	expect := 1
	runSample(t, root, expect)
}

func TestSample2(t *testing.T) {
	root := new(TreeNode)
	root.Left = new(TreeNode)
	root.Left.Left = new(TreeNode)
	root.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Right = new(TreeNode)
	expect := 2
	runSample(t, root, expect)
}

func TestSample3(t *testing.T) {
	root := new(TreeNode)
	expect := 1
	runSample(t, root, expect)
}

func TestSample4(t *testing.T) {
	root := new(TreeNode)
	root.Left = new(TreeNode)
	root.Left.Left = new(TreeNode)
	root.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Right = new(TreeNode)
	root.Left.Left.Left.Right.Right = new(TreeNode)
	expect := 2
	runSample(t, root, expect)
}
