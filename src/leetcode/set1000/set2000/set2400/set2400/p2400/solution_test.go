package p2400

import "testing"

func runSample(t *testing.T, start, end int, k int, expect int) {
	res := numberOfWays(start, end, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	start, end := 1, 2
	k := 3
	expect := 3
	runSample(t, start, end, k, expect)
}

func TestSample2(t *testing.T) {
	start, end := 2, 5
	k := 10
	expect := 0
	runSample(t, start, end, k, expect)
}

func TestSample3(t *testing.T) {
	start, end := 1, 100
	k := 891
	expect := 477741651
	runSample(t, start, end, k, expect)
}
