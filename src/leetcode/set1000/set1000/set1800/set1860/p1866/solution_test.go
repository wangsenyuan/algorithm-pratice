package p1866

import "testing"

func runSample(t *testing.T, n, k int, expect int) {
	res := rearrangeSticks(n, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	expect := 3
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 5
	expect := 1
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 20, 11
	expect := 647427950
	runSample(t, n, k, expect)
}
