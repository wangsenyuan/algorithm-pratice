package p2318

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := distinctSequences(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 184)
}


func TestSample2(t *testing.T) {
	runSample(t, 2, 22)
}
