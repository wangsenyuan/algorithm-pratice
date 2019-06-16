package p236

import "testing"

func runSample(t *testing.T, root *TreeNode, p, q *TreeNode, expect int) {
	res := lowestCommonAncestor(root, p, q)

	if res.Val != expect {
		t.Error("Sample wrong")
	}
}

func sample() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	return root
}

func TestSample1(t *testing.T) {
	p := &TreeNode{Val: 5}
	q := &TreeNode{Val: 1}
	runSample(t, sample(), p, q, 3)
}

func TestSample2(t *testing.T) {
	p := &TreeNode{Val: 5}
	q := &TreeNode{Val: 4}
	runSample(t, sample(), p, q, 5)
}

func TestSample3(t *testing.T) {
	p := &TreeNode{Val: 5}
	q := &TreeNode{Val: 8}
	runSample(t, sample(), p, q, 3)
}

func TestSample4(t *testing.T) {
	p := &TreeNode{Val: 4}
	q := &TreeNode{Val: 8}
	runSample(t, sample(), p, q, 3)
}

func TestSample5(t *testing.T) {
	p := &TreeNode{Val: 4}
	q := &TreeNode{Val: 6}
	runSample(t, sample(), p, q, 5)
}

func TestSample6(t *testing.T) {
	p := &TreeNode{Val: 7}
	q := &TreeNode{Val: 1}
	runSample(t, sample(), p, q, 3)
}
