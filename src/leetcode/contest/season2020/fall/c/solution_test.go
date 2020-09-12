package c

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minimumOperations(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "rrryyyrryyyrr", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, "ryr", 0)
}
