package p2561

import "testing"

func runSample(t *testing.T, pos []int, k int, expect int) {
	res := maximizeWin(pos, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pos := []int{1, 1, 2, 2, 3, 3, 5}
	k := 2
	expect := 7
	runSample(t, pos, k, expect)
}

func TestSample2(t *testing.T) {
	pos := []int{1, 2, 3, 4}
	k := 0
	expect := 2
	runSample(t, pos, k, expect)
}
