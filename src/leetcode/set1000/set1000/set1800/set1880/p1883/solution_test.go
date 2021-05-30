package p1883

import "testing"

func runSample(t *testing.T, n string, x int, expect string) {
	res := maxValue(n, x)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := "99"
	x := 9
	expect := "999"
	runSample(t, n, x, expect)
}

func TestSample2(t *testing.T) {
	n := "-13"
	x := 2
	expect := "-123"
	runSample(t, n, x, expect)
}
