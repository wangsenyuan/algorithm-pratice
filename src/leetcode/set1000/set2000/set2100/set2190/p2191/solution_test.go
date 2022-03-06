package p2191

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minMovesToMakePalindrome(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aabb"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "letelt"
	expect := 2
	runSample(t, s, expect)
}
