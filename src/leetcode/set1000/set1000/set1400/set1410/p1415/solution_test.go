package p1415

import "testing"

func runSample(t *testing.T, n, k int, expect string) {
	res := getHappyString(n, k)

	if res != expect {
		t.Errorf("Sample %d %d, expect %s, but got %s", n, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 3, "c")
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 9, "cab")
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 100, "abacbabacb")
}
