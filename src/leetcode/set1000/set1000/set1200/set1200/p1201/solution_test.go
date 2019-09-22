package p1201

import "testing"

func runSample(t *testing.T, n int, a, b, c int, expect int) {
	res := nthUglyNumber(n, a, b, c)
	if res != expect {
		t.Errorf("Sample %d %d %d %d, expect %d, but got %d", n, a, b, c, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	a := 2
	b := 3
	c := 5
	runSample(t, n, a, b, c, 4)
}

func TestSample2(t *testing.T) {
	n := 4
	a := 2
	b := 3
	c := 4
	expect := 6
	runSample(t, n, a, b, c, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	a := 2
	b := 11
	c := 13
	expect := 10
	runSample(t, n, a, b, c, expect)
}

func TestSample4(t *testing.T) {
	n := 1000000000
	a := 2
	b := 217983653
	c := 336916467
	expect := 1999999984
	runSample(t, n, a, b, c, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	a := 2
	b := 3
	c := 3
	expect := 8
	runSample(t, n, a, b, c, expect)
}
