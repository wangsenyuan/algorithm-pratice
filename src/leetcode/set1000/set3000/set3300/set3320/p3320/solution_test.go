package p3320

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := countWinningSequences(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "FFF"
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "FWEFW"
	expect := 18
	runSample(t, s, expect)
}