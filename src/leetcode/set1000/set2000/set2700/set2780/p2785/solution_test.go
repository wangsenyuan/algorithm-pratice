package p2785

import "testing"

func runSample(t *testing.T, n int, x int, expect int) {
	res := numberOfWays(n, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	x := 2
	expect := 1
	runSample(t, n, x, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	x := 1
	expect := 2
	runSample(t, n, x, expect)
}

func TestSample3(t *testing.T) {
	n := 1
	x := 1
	expect := 1
	runSample(t, n, x, expect)
}

func TestSample4(t *testing.T) {
	n := 18
	x := 1
	expect := 46
	runSample(t, n, x, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	x := 2
	expect := 1
	runSample(t, n, x, expect)
}
