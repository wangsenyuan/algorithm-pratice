package p1419

import "testing"

func runSample(t *testing.T, str string, expect int) {
	res := minNumberOfFrogs(str)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", str, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "croakcroak", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "crcoakroak", 2)
}

func TestSample3(t *testing.T) {
	runSample(t, "croakcrook", -1)
}

func TestSample4(t *testing.T) {
	runSample(t, "croakcroa", -1)
}
