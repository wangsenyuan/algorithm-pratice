package p3504

import "testing"

func runSample(t *testing.T, s string, x string, expect int) {
	res := longestPalindrome(s, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "a"
	x := "a"
	expect := 2
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "abc"
	x := "def"
	expect := 1
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "b"
	x := "aaaa"
	expect := 4
	runSample(t, s, x, expect)
}

func TestSample4(t *testing.T) {
	s := "abcde"
	x := "ecdba"
	expect := 5
	runSample(t, s, x, expect)
}

func TestSample5(t *testing.T) {
	s := "i"
	x := "ih"
	expect := 2
	runSample(t, s, x, expect)
}

func TestSample6(t *testing.T) {
	s := "axchh"
	x := "semjc"
	// chhc
	expect := 4
	runSample(t, s, x, expect)
}
