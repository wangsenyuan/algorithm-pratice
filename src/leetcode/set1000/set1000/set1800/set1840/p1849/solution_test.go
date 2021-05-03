package p1849

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := splitString(s)

	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0090089"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "050043"
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "10009998"
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "9080701"
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "10"
	expect := true
	runSample(t, s, expect)
}
