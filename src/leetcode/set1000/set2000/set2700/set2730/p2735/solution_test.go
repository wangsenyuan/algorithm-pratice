package p2735

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := smallestString(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "cbabc"
	expect := "baabc"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "a"
	expect := "z"
	runSample(t, s, expect)
}
