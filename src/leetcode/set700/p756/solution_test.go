package p756

import "testing"

func runSample(t *testing.T, bottom string, allowed []string, expect bool) {
	res := pyramidTransition(bottom, allowed)
	if res != expect {
		t.Errorf("sample: %s %v; should give answer %v, but got %v", bottom, allowed, expect, res)
	}
}

func TestSample1(t *testing.T) {
	bottom := "XYZ"
	allowed := []string{
		"XYD", "YZE", "DEA", "FFF",
	}
	expect := true
	runSample(t, bottom, allowed, expect)
}

func TestSample2(t *testing.T) {
	bottom := "XXYX"
	allowed := []string{
		"XXX", "XXY", "XYX", "XYY", "YXZ",
	}
	expect := false
	runSample(t, bottom, allowed, expect)
}
