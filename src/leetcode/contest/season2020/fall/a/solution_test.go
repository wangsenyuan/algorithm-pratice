package a

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := calculate(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "AB", 4)
}
