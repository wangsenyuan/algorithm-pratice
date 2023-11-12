package p2929

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := stringCount(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 12)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 83943898)
}
