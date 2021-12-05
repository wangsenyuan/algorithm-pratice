package p2095

import "testing"

func runSample(t *testing.T, root *TreeNode, start, dest int, expect string) {
	res := getDirections(root, start, dest)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 4}
	start := 3
	dest := 6
	expect := "UURL"
	runSample(t, root, start, dest, expect)
}
