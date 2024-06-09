package p3178

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := valueAfterKSeconds(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 5, 56)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 3, 35)
}
