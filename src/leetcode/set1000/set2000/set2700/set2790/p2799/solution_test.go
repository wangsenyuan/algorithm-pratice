package p2799

import "testing"

func runSample(t *testing.T, a, b, c string, expect string) {
	res := minimumString(a, b, c)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "abc"
	b := "bca"
	c := "aaa"
	expect := "aaabca"
	runSample(t, a, b, c, expect)
}

func TestSample2(t *testing.T) {
	a := "ab"
	b := "ba"
	c := "aba"
	expect := "aba"
	runSample(t, a, b, c, expect)
}

func TestSample3(t *testing.T) {
	a := "ca"
	b := "a"
	c := "a"
	expect := "ca"
	runSample(t, a, b, c, expect)
}

func TestSample4(t *testing.T) {
	a := "b"
	b := "a"
	c := "aba"
	expect := "aba"
	runSample(t, a, b, c, expect)
}
