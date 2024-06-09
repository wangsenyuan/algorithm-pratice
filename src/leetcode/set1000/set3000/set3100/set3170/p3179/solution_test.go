package p3179

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := numberOfChild(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 6, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 23, 43, 1)
}
