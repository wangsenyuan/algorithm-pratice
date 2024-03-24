package p3091

import "testing"

func runSample(t *testing.T, k int, expect int) {
	res := minOperations(k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 11, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 0)
}
