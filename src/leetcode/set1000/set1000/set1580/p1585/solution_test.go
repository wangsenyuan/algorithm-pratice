package p1585

import "testing"

func runSample(tc *testing.T, s, t string, expect bool) {
	res := isTransformable(s, t)

	if res != expect {
		tc.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(tc *testing.T) {
	s := "84532"
	t := "34852"
	expect := true
	runSample(tc, s, t, expect)
}

func TestSample2(tc *testing.T) {
	s := "34521"
	t := "23415"
	expect := true
	runSample(tc, s, t, expect)
}

func TestSample3(tc *testing.T) {
	s := "12345"
	t := "12435"
	expect := false
	runSample(tc, s, t, expect)
}

func TestSample4(tc *testing.T) {
	s := "1"
	t := "2"
	expect := false
	runSample(tc, s, t, expect)
}

func TestSample5(tc *testing.T) {
	s := "2032"
	t := "0223"
	expect := true
	runSample(tc, s, t, expect)
}

func TestSample6(tc *testing.T) {
	s := "2032"
	t := "0223"
	expect := true
	runSample(tc, s, t, expect)
}

func TestSample7(tc *testing.T) {
	s := "432513576"
	t := "231435567"
	expect := true
	runSample(tc, s, t, expect)
}
