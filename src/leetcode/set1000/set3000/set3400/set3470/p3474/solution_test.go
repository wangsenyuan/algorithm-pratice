package p3474

import (
	"testing"
)

func runSample(t *testing.T, s1 string, s2 string, expect string) {
	res := generateString(s1, s2)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s1 := "TFTF"
	s2 := "ab"
	expect := "ababa"
	runSample(t, s1, s2, expect)
}

func TestSample2(t *testing.T) {
	s1 := "TFTF"
	s2 := "abc"
	expect := ""
	runSample(t, s1, s2, expect)
}

func TestSample3(t *testing.T) {
	s1 := "F"
	s2 := "d"
	expect := "a"
	runSample(t, s1, s2, expect)
}

func TestSample4(t *testing.T) {
	s1 := "FFTFFF"
	s2 := "a"
	expect := "bbabbb"
	runSample(t, s1, s2, expect)
}
