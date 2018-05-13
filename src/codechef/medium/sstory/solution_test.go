package main

import "testing"

func runSample1(t *testing.T, s1, s2 string, expect string) {
	length, begin := solve([]byte(s1), []byte(s2))

	if len(expect) != length {
		t.Errorf("Sample %s %s, expect %s, but got length %d", s1, s2, expect, length)
	} else if len(expect) > 0 && string(s2[begin:begin+length]) != expect {
		t.Errorf("Sample %s %s, expect %s, but got %s", s1, s2, expect, string(s2[begin:begin+length]))
	}
}

func runSample(t *testing.T, s1, s2 string, expect string) {
	res := solve3([]byte(s1), []byte(s2))

	if res != expect {
		t.Errorf("Sample %s %s, expect %s, but got %s", s1, s2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s1 := "adsyufsfdsfdsf"
	s2 := "fdyusgfdfyu"
	expect := "fd"
	runSample(t, s1, s2, expect)
}

func TestSample2(t *testing.T) {
	s1 := "adsyufsfdsfdsf"
	s2 := "bbbbbbbbb"
	expect := ""
	runSample(t, s1, s2, expect)
}

func TestSample3(t *testing.T) {
	s1 := "abcdef"
	s2 := "defabc"
	expect := "def"
	runSample(t, s1, s2, expect)
}

func TestSample4(t *testing.T) {
	s1 := "abcddffed"
	s2 := "afdfedefabc"
	expect := "fed"
	runSample(t, s1, s2, expect)
}
