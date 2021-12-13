package p2028

import "testing"

func runSample(t *testing.T, stones []int, expect bool) {
	res := stoneGameIX(stones)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	stones := []int{5, 1, 2, 4, 3}
	expect := false
	runSample(t, stones, expect)
}

func TestSample2(t *testing.T) {
	stones := []int{1, 2}
	expect := true
	runSample(t, stones, expect)
}
