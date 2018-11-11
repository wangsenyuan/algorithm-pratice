package p938

import "testing"

func runSample(t *testing.T, root *TreeNode, L, R int, expect int) {
	res := rangeSumBST(root, L, R)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	root := ParseAsTree("[10,5,15,3,7,null,18]")
	L, R := 7, 15
	expect := 32
	runSample(t, root, L, R, expect)
}

func TestSample2(t *testing.T) {
	root := ParseAsTree("[10,5,15,3,7,13,18,1,null,6]")
	L, R := 6, 10
	expect := 23
	runSample(t, root, L, R, expect)
}

// [18,9,27,6,15,24,30,3,null,12,null,21]
// 18
// 24
// 63

func TestSample3(t *testing.T) {
	root := ParseAsTree("[18,9,27,6,15,24,30,3,null,12,null,21]")
	L, R := 18, 24
	expect := 63
	runSample(t, root, L, R, expect)
}
