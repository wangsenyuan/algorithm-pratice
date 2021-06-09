package p1888

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minFlips(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "111000"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "010"
	expect := 0
	runSample(t, s, expect)
}
