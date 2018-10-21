package p925

import "testing"

func runSample(t *testing.T, name, typed string, expect bool) {
	res := isLongPressedName(name, typed)

	if res != expect {
		t.Errorf("sample %s %s, expect %t, but got %t", name, typed, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "alex", "aaleex", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "saeed", "ssaaedd", false)
}

func TestSample3(t *testing.T) {
	runSample(t, "leelee", "lleeelee", true)
}

func TestSample4(t *testing.T) {
	runSample(t, "laiden", "laiden", true)
}

func TestSample5(t *testing.T) {
	runSample(t, "pyplrz", "ppyypllr", false)
}
