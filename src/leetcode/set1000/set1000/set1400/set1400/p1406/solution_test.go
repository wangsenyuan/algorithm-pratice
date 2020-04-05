package p1406

import "testing"

func runSample(t *testing.T, stones []int, expect string) {
	res := stoneGameIII(stones)

	if res != expect {
		t.Errorf("Sample %v, expect %s, but got %s", stones, expect, res)
	}
}

func TestSample1(t *testing.T) {
	stones := []int{1, 2, 3, 7}
	expect := "Bob"
	runSample(t, stones, expect)
}

func TestSample2(t *testing.T) {
	stones := []int{1, 2, 3, -9}
	expect := "Alice"
	runSample(t, stones, expect)
}

func TestSample3(t *testing.T) {
	stones := []int{1, 2, 3, 6}
	expect := "Tie"
	runSample(t, stones, expect)
}

func TestSample4(t *testing.T) {
	stones := []int{1, 2, 3, -1, -2, -3, 7}
	expect := "Alice"
	runSample(t, stones, expect)
}
