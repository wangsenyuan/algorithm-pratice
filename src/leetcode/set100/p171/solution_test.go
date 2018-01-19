package p171

import "testing"

func TestSample1(t *testing.T) {
	s := "AA"
	expect := 27
	res := titleToNumber(s)
	if res != expect {
		t.Errorf("sample %s should give %d, but got %d", s, expect, res)
	}
}

func TestSample2(t *testing.T) {
	s := "AB"
	expect := 28
	res := titleToNumber(s)
	if res != expect {
		t.Errorf("sample %s should give %d, but got %d", s, expect, res)
	}
}

func TestSample3(t *testing.T) {
	s := "Z"
	expect := 26
	res := titleToNumber(s)
	if res != expect {
		t.Errorf("sample %s should give %d, but got %d", s, expect, res)
	}
}
