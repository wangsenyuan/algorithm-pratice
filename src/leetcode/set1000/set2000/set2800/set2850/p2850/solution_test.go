package p2850

import "testing"

func runSample(t *testing.T, sx, sy, fx, fy int, time int, expect bool) {
	res := isReachableAtTime(sx, sy, fx, fy, time)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	sx, sy, fx, fy := 1, 1, 1, 1
	time := 3
	expect := true
	runSample(t, sx, sy, fx, fy, time, expect)
}
