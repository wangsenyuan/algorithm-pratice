package p1604

import "testing"

func runSample(t *testing.T, a, b string, expect bool) {
	res := checkPalindromeFormation(a, b)

	if res != expect {
		t.Errorf("Sample %s %s, expect %t, but got %t", a, b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "x"
	b := "y"
	expect := true
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "ulacfd"
	b := "jizalu"
	expect := true
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := "abxyxdd"
	b := "12344ba"
	expect := true
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := "pvhmupgqeltozftlmfjjde"
	b := "yjgpzbezspnnpszebzmhvp"
	expect := true
	runSample(t, a, b, expect)
}
