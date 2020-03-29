package p1396

import "testing"

func runSample(t *testing.T, n int, s1, s2 string, evil string, expect int) {
	res := findGoodStrings(n, s1, s2, evil)

	if res != expect {
		t.Errorf("Sample %d %s %s %s, expect %d, but got %d", n, s1, s2, evil, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	s1 := "aa"
	s2 := "da"
	evil := "b"
	expect := 51
	runSample(t, n, s1, s2, evil, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	s1 := "leetcode"
	s2 := "leetgoes"
	evil := "leet"
	expect := 0
	runSample(t, n, s1, s2, evil, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	s1 := "gx"
	s2 := "gz"
	evil := "x"
	expect := 2
	runSample(t, n, s1, s2, evil, expect)
}

func TestSample4(t *testing.T) {
	n := 8
	s1 := "pzdanyao"
	s2 := "wgpmtywi"
	evil := "sdka"
	expect := 500543753
	runSample(t, n, s1, s2, evil, expect)
}
