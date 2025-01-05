package p3407

import "testing"

func runSample(t *testing.T, s string, p string, expect bool) {
	res := hasMatch(s, p)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "xks"
	p := "s*"
	expect := true
	runSample(t, s, p, expect)
}

func TestSample2(t *testing.T) {
	s := "car"
	p := "c*v"
	expect := false
	runSample(t, s, p, expect)
}
