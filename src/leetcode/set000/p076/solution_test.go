package p076

import "testing"

func TestSample(t *testing.T) {
	s := "ADOBECODEBANC"
	x := "ABC"
	expect := "BANC"
	res := minWindow(s, x)
	if res != expect {
		t.Errorf("sample %s %s, expect %s, but got %s", s, x, expect, res)
	}
}

func TestSample2(t *testing.T) {
	s := "bba"
	x := "ab"
	expect := "ba"
	res := minWindow(s, x)
	if res != expect {
		t.Errorf("sample %s %s, expect %s, but got %s", s, x, expect, res)
	}
}

func TestSample3(t *testing.T) {
	s := "a"
	x := "a"
	expect := "a"
	res := minWindow(s, x)
	if res != expect {
		t.Errorf("sample %s %s, expect %s, but got %s", s, x, expect, res)
	}
}
