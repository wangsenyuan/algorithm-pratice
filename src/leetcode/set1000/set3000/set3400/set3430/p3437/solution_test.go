package p3437

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := maxDistance(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "NWSE"
	k := 1
	expect := 3
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "NSWWEW"
	k := 3
	expect := 6
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "EE"
	k := 1
	expect := 2
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "NWSE"
	k := 1
	expect := 3
	runSample(t, s, k, expect)
}

func TestSample5(t *testing.T) {
	s := "NEEW"
	k := 1
	expect := 4
	runSample(t, s, k, expect)
}

func TestSample6(t *testing.T) {
	s := "SN"
	k := 0
	expect := 1
	runSample(t, s, k, expect)
}
