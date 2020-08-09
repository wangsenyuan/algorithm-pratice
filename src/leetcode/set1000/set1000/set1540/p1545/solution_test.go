package p1545

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := longestAwesome(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "3242415"
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "00"
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "373781"
	expect := 5
	runSample(t, s, expect)
}
