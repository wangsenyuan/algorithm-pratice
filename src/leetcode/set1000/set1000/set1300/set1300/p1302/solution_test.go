package p1302

import "testing"

func runSample(t *testing.T, root *TreeNode, expect int) {
	res := deepestLeavesSum(root)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	root := new(TreeNode)
	root.Val = 1
	expect := 1
	runSample(t, root, expect)
}
