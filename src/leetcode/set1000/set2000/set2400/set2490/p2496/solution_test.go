package p2496

import "testing"

func runSample(t *testing.T, stones []int, expect int) {
	res := maxJump(stones)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	stones := []int{0, 2, 5, 6, 7}
	expect := 5
	runSample(t, stones, expect)
}

func TestSample2(t *testing.T) {
	stones := []int{0, 3, 9}
	expect := 9
	runSample(t, stones, expect)
}

func TestSample3(t *testing.T) {
	stones := []int{0, 5, 12, 25, 28, 35}
	expect := 20
	runSample(t, stones, expect)
}
