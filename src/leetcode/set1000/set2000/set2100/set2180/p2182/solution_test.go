package p2182

import "testing"

func runSample(t *testing.T, s string, rep int, expect string) {
	res := repeatLimitedString(s, rep)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "cczazcc"
	rep := 3
	expect := "zzcccac"
	runSample(t, s, rep, expect)
}


func TestSample2(t *testing.T) {
	s := "aababab"
	rep := 2
	expect := "bbabaa"
	runSample(t, s, rep, expect)
}
