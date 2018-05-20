package p836

import "testing"

func runSample(t *testing.T, rec1 []int, rec2 []int, expect bool) {
	res := isRectangleOverlap(rec1, rec2)
	if res != expect {
		t.Errorf("sample %v %v, expect %t, but got %t", rec1, rec2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	rec1 := []int{0, 0, 2, 2}
	rec2 := []int{1, 1, 3, 3}
	runSample(t, rec1, rec2, true)
}

func TestSample2(t *testing.T) {
	rec1 := []int{0, 0, 1, 1}
	rec2 := []int{1, 0, 2, 1}
	runSample(t, rec1, rec2, false)
}
