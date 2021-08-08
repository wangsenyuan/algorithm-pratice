package p1958

import "testing"

func runSample(t *testing.T, s string, expect int64) {
	res := maxProduct(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ababbb"
	var expect int64 = 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "zaaaxbbby"
	var expect int64 = 9
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "ggbswiymmlevedhkbdhntnhdbkhdevelmmyiwsbgg"
	var expect int64 = 45
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "aaaaaa"
	var expect int64 = 9
	runSample(t, s, expect)
}
