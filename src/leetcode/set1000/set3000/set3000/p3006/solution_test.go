package p3006

import "testing"

func runSample(t *testing.T, k int64, x int, expect int64) {
	res := findMaximumNumber(k, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 9
	x := 1
	expect := 6
	runSample(t, int64(k), x, int64(expect))
}

func TestSample2(t *testing.T) {
	k := 7
	x := 2
	expect := 9
	runSample(t, int64(k), x, int64(expect))
}
