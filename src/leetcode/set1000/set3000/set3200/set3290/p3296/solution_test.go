package p3296

import (
	"testing"
)

func runSample(t *testing.T, h int, w []int, expect int) {
	res := minNumberOfSeconds(h, w)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	h := 10
	w := []int{3, 2, 2, 4}
	expect := 12
	runSample(t, h, w, expect)
}

func TestSample2(t *testing.T) {
	h := 5
	w := []int{1}
	expect := 15
	runSample(t, h, w, expect)
}
