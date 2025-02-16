package p3455

import (
	"testing"
)

func runSample(t *testing.T, s string, p string, expect int) {
	res := shortestMatchingSubstring(s, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abaacbaecebce"
	p := "ba*c*ce"
	expect := 8
	runSample(t, s, p, expect)
}

func TestSample2(t *testing.T) {
	s := "baccbaadbc"
	p := "cc*baa*adb"
	expect := -1
	runSample(t, s, p, expect)
}

func TestSample3(t *testing.T) {
	s := "a"
	p := "**"
	expect := 0
	runSample(t, s, p, expect)
}

func TestSample4(t *testing.T) {
	s := "madlogic"
	p := "*adlogi*"
	expect := 6
	runSample(t, s, p, expect)
}
