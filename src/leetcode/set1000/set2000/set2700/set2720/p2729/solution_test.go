package p2729

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := longestSemiRepetitiveSubstring(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "52233"
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "5494"
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "1111111"
	expect := 2
	runSample(t, s, expect)
}
