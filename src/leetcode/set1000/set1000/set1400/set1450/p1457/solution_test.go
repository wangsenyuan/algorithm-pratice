package p1457

import "testing"

func TestSample1(t *testing.T) {
	node := new(TreeNode)
	node.Val = 1

	res := pseudoPalindromicPaths(node)

	if res != 1 {
		t.Errorf("Sample expect 1, but got %d", res)
	}
}
