package p951

import "testing"

func runSample(t *testing.T, root1, root2 string, expect bool) {
	tree1 := parseNode(root1)
	tree2 := parseNode(root2)
	res := flipEquiv(tree1, tree2)

	if res != expect {
		t.Errorf("Sample %s %s expect %t, but got %t", root1, root2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "[1,2,3,4,5,6,null,null,null,7,8]", "[1,3,2,null,6,4,5,null,null,null,null,8,7]", true)
}
