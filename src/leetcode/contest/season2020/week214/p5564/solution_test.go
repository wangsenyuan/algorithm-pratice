package p5564

import "testing"

func runSample(t *testing.T, insts []int, expect int) {
	res := createSortedArray(insts)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", insts, expect, res)
	}
}

func TestSample1(t *testing.T) {
	insts := []int{1, 5, 6, 2}
	expect := 1
	runSample(t, insts, expect)
}

func TestSample2(t *testing.T) {
	insts := []int{1, 2, 3, 6, 5, 4}
	expect := 3
	runSample(t, insts, expect)
}

func TestSample3(t *testing.T) {
	insts := []int{1, 3, 3, 3, 2, 4, 2, 1, 2}
	expect := 4
	runSample(t, insts, expect)
}


