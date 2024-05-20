package p3151

import "testing"

func runSample(t *testing.T, k int, expect int) {
	res := waysToReachStair(k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	expect := 4
	runSample(t, k, expect)
}

func TestSample2(t *testing.T) {
	k := 1020
	expect := 330
	runSample(t, k, expect)
}
