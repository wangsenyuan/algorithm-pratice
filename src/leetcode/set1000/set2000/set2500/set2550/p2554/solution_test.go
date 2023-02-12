package p2554

import "testing"

func runSample(t *testing.T, s string, x string, expect int) {
	res := minimumScore(s, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abacaba"
	x := "bzaa"
	expect := 1
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "cde"
	x := "xyz"
	expect := 3
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "acdedcdbabecdbebda"
	x := "bbecddb"
	expect := 1
	runSample(t, s, x, expect)
}

func TestSample4(t *testing.T) {
	s := "eeecaeecdeeadcdbcaa"
	x := "edecabe"
	expect := 2
	runSample(t, s, x, expect)
}

func TestSample5(t *testing.T) {
	s := "eceecbabe"
	x := "bdeaec"
	expect := 4
	runSample(t, s, x, expect)
}
