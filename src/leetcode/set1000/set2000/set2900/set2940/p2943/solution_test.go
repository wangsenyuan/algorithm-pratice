package p2943

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := minimumCoins(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 10, 1, 1}
	expect := 2
	runSample(t, a, expect)
}
