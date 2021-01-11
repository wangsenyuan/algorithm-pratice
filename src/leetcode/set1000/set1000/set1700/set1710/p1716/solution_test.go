package p1716

import "testing"

func runSample(t *testing.T, s string, x, y int, expect int) {
	res := maximumGain(s, x, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "cdbcbbaaabab"
	x := 4
	y := 5
	expect := 19
	runSample(t, s, x, y, expect)
}

func TestSample2(t *testing.T) {
	s := "aabbaaxybbaabb"
	x := 5
	y := 4
	expect := 20
	runSample(t, s, x, y, expect)
}

func TestSample3(t *testing.T) {
	s := "aabbabkbbbfvybssbtaobaaaabataaadabbbmakgabbaoapbbbbobaabvqhbbzbbkapabaavbbeghacabamdpaaqbqabbjbababmbakbaabajabasaabbwabrbbaabbafubayaazbbbaababbaaha"
	x := 1926
	y := 4320
	expect := 112374
	runSample(t, s, x, y, expect)
}
