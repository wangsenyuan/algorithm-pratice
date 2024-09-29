package p3303

import (
	"testing"
)

func runSample(t *testing.T, s string, p string, expect int) {
	res := minStartingIndex(s, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcdefg"
	pattern := "bcdffg"
	expect := 1
	runSample(t, s, pattern, expect)
}

func TestSample2(t *testing.T) {
	s := "ababbababa"
	pattern := "bacaba"
	expect := 4
	runSample(t, s, pattern, expect)
}

func TestSample3(t *testing.T) {
	s := "abcd"
	pattern := "dba"
	expect := -1
	runSample(t, s, pattern, expect)
}

func TestSample4(t *testing.T) {
	s := "dde"
	pattern := "d"
	expect := 0
	runSample(t, s, pattern, expect)
}
