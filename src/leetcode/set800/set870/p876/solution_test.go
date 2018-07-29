package p876

import "testing"

func runSample(t *testing.T, s string, expect int) {
	head := ParseAsList(s)
	res := middleNode(head)
	if res.Val != expect {
		t.Errorf("Sample %s, expect %d, but got %v", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "[1,2,3,4,5]", 3)
}

func TestSample2(t *testing.T) {
	runSample(t, "[1,2,3,4,5,6]", 4)
}

func TestSample3(t *testing.T) {
	runSample(t, "[1]", 1)
}
