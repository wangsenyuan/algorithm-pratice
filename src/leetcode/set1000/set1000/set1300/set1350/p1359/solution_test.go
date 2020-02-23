package p1359

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := countOrders(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 90)
}
