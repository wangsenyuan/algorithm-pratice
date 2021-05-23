package p1871

import "testing"

func runSample(t *testing.T, s string, min, max int, expect bool) {
	res := canReach(s, min, max)

	if res != expect {
		t.Errorf("Smaple expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "011010"
	min := 2
	max := 3
	expect := true

	runSample(t, s, min, max, expect)
}

func TestSample2(t *testing.T) {
	s := "01101110"
	min := 2
	max := 3
	expect := false

	runSample(t, s, min, max, expect)
}

func TestSample3(t *testing.T) {
	s := "00111010"
	min := 3
	max := 5
	expect := false

	runSample(t, s, min, max, expect)
}
