package p2167

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minimumTime(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1100101"
	expect := 5
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := "0010"
	expect := 2
	runSample(t, s, expect)
}
