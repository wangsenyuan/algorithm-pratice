package lcp43

import "testing"

func runSample(t *testing.T, dirs []string, expect int) {
	res := trafficCommand(dirs)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	dirs := []string{"W", "N", "ES", "W"}
	expect := 2
	runSample(t, dirs, expect)
}

func TestSample2(t *testing.T) {
	dirs := []string{"NS", "WE", "SE", "EW"}
	expect := 3
	runSample(t, dirs, expect)
}

func TestSample3(t *testing.T) {
	dirs := []string{"SSS", "WWW", "SSSSSSS", "SSSSSSSSSSSS"}
	expect := 22
	runSample(t, dirs, expect)
}
