package p1071

import "testing"

func runSample(t *testing.T, str1 string, str2 string, expect string) {
	res := gcdOfStrings(str1, str2)
	if res != expect {
		t.Errorf("Sample %s %s, expect %s, but got %s", str1, str2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	str1 := "TAUXXTAUXXTAUXXTAUXXTAUXX"
	str2 := "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"
	expect := "TAUXX"
	runSample(t, str1, str2, expect)
}
