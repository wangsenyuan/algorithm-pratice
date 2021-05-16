package p1864

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minSwaps(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "111000"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "010"
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "1110"
	expect := -1
	runSample(t, s, expect)
}
