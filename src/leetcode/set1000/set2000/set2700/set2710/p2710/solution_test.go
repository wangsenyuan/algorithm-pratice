package p2710

import "testing"

func runSample(t *testing.T, s string, expect int64) {
	res := minimumCost(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0011"
	var expect int64 = 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "010101"
	var expect int64 = 9
	runSample(t, s, expect)
}
