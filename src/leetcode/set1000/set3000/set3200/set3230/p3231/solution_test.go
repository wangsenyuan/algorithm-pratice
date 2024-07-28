package p3231

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := numberOfSubstrings(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "00011"
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "101101"
	expect := 16
	runSample(t, s, expect)
}
