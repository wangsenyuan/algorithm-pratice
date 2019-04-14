package p1031

import "testing"

func sameTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && sameTree(a.Left, b.Left) && sameTree(a.Right, b.Right)
}

func runSample(t *testing.T, S string, expect *TreeNode) {
	res := recoverFromPreorder(S)

	if !sameTree(res, expect) {
		t.Errorf("Sample %s, get wrong answer", S)
	}
}

func TestSample1(t *testing.T) {
	S := "1-2--3--4-5--6--7"
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 3
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 4
	root.Right = new(TreeNode)
	root.Right.Val = 5
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 6
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 7
	runSample(t, S, root)
}
