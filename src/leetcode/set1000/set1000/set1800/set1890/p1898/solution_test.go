package p1898

import "testing"

func runSample(t *testing.T, s, p string, r []int, expect int) {
	res := maximumRemovals(s, p, r)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcacb"
	p := "ab"
	r := []int{3, 1, 0}
	expect := 2
	runSample(t, s, p, r, expect)
}

func TestSample2(t *testing.T) {
	s := "abcbddddd"
	p := "abcd"
	r := []int{3, 2, 1, 4, 5, 6}
	expect := 1
	runSample(t, s, p, r, expect)
}

func TestSample3(t *testing.T) {
	s := "abcab"
	p := "abc"
	r := []int{0, 1, 2, 3, 4}
	expect := 0
	runSample(t, s, p, r, expect)
}
