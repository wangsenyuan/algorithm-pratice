package main

import "testing"

func runSample(t *testing.T, ord []string, expect string) {
	res := solve(ord)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ord := []string{
		"NAUTILUS",
		"FIREFOX",
		"GEDIT",
		"FIREFOX",
		"NAUTILUS",
		"WINE",
		"GIMP",
		"TERMINAL",
		"WINE",
	}
	expect := "NEALMPUSOXIT"
	runSample(t, ord, expect)
}

func TestSample2(t *testing.T) {
	ord := []string{
		"YDK",
		"IDK",
	}
	expect := "DKDK"
	runSample(t, ord, expect)
}
