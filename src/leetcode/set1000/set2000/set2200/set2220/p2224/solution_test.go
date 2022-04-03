package p2224

import "testing"

func runSample(t *testing.T, current string, correct string, expect int) {
	res := convertTime(current, correct)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "02:30"
	b := "04:35"
	expect := 3
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "11:00"
	b := "11:01"
	expect := 1
	runSample(t, a, b, expect)
}
