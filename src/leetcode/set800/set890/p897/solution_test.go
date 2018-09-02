package p897

import "testing"

func runSample(t *testing.T, s string, expect string) {
	root := ParseAsTree(s)
	root = increasingBST(root)

	res := SprintTree(root)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "[5,3,6,2,4,null,8,1,null,null,null,7,9]", "[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9,null,null]")
}
