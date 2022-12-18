package p2505

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := smallestValue(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 15, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 3)
}


func TestSample3(t *testing.T) {
	runSample(t, 4, 4)
}
