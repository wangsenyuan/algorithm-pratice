package p5606

import "testing"

func runSample(t *testing.T, n, k int, expect string) {
	res := getSmallestString(n, k)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 27
	expect := "aay"
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 73
	expect := "aaszz"
	runSample(t, n, k, expect)
}
