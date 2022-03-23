package p2212

import "testing"

func runSample(t *testing.T, numArrows int, alice []int, expect []int) {
	res := maximumBobPoints(numArrows, alice)
	expectScore := checkWin(alice, expect)
	resScore := checkWin(alice, res)

	if expectScore != resScore {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func checkWin(alice, bob []int) int {
	var res int
	for i := 0; i < 12; i++ {
		if alice[i] < bob[i] {
			res += i
		}
	}
	return res
}

func TestSample1(t *testing.T) {
	numArrows := 9
	alice := []int{1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2, 0}
	expect := []int{0, 0, 0, 0, 1, 1, 0, 0, 1, 2, 3, 1}
	runSample(t, numArrows, alice, expect)
}

func TestSample2(t *testing.T) {
	numArrows := 3
	alice := []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 2}
	expect := []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0}
	runSample(t, numArrows, alice, expect)
}
