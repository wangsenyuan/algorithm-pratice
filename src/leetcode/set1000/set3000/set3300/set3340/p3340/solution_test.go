package p3340

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := isBalanced(s)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1234"
	expect := false
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "24123"
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "11"
	expect := true
	runSample(t, s, expect)
}
