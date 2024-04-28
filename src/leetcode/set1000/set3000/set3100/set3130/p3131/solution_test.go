package p3131

import "testing"

func runSample(t *testing.T, n int, x int, expect int) {
	res := minEnd(n, x)

	if int(res) != expect {
		t.Fatalf("Sample %d %d, expect %d, but got %d", n, x, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	x := 4
	expect := 6
	runSample(t, n, x, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	x := 7
	expect := 15
	runSample(t, n, x, expect)
}

func TestSample3(t *testing.T) {
	n := 10000000
	x := 39
	expect := 159999999
	runSample(t, n, x, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	x := 12

	// 1100, 1101, 1110, 1111
	// 11100  11101 11110 11111

	expect := 29
	runSample(t, n, x, expect)
}
