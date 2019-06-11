package p1081

import "testing"

func runSample(t *testing.T, s string, limit int, expect string) {
	root := ParseAsTree(s)
	res := sufficientSubset(root, limit)

	rs := SprintTree(res)

	if rs != expect {
		t.Errorf("Sample %s %d, expect %s, but got %s", s, limit, expect, rs)
	}
}

func TestSample1(t *testing.T) {
	tree := "[5,4,8,11,null,17,4,7,1,null,null,5,3]"
	limit := 22
	expect := "[5,4,8,11,null,17,4,7,null,null,null,5]"
	runSample(t, tree, limit, expect)
}
