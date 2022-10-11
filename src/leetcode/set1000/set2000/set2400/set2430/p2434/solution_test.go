package p2434

import "testing"


func runSample(t *testing.T, s string, expect string) {
	res := robotWithString(s)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "zza"
	expect := "azz"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "bdda"
	expect := "addb"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "bac"
	expect := "abc"
	runSample(t, s, expect)
}