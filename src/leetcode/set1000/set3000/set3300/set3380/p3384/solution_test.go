package p3384

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := makeStringGood(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	s := "acab"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "wddw"
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "aaabc"
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "ruuu"
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "gigigjjggjjgg"
	expect := 3
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "cabbcbacbbbbbaa"
	expect := 4
	runSample(t, s, expect)
}
