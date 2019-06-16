package p867

import "testing"

func runSample(t *testing.T, str string, expect int) {
	tree := ParseStrAsTree(str)
	res := subtreeWithAllDeepest(tree)

	if res.Val != expect {
		t.Errorf("Sample %s, expect %d, but got %d", str, expect, res.Val)
	}
}

func TestSample1(t *testing.T) {
	str := "[3,5,1,6,2,0,8,null,null,7,4]"
	expect := 2
	runSample(t, str, expect)
}

func TestSample2(t *testing.T) {
	str := "[1]"
	expect := 1
	runSample(t, str, expect)
}

func TestSample3(t *testing.T) {
	str := "[1,2,null]"
	expect := 2
	runSample(t, str, expect)
}

func TestSample4(t *testing.T) {
	str := "[0,1,3,null,2]"
	expect := 2
	runSample(t, str, expect)
}
