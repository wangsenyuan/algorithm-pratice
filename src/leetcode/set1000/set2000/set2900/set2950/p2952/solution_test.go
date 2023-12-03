package p2952

import "testing"

func runSample(t *testing.T, n int, sick []int, expect int) {
	res := numberOfSequence(n, sick)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	sick := []int{0, 4}
	expect := 4
	runSample(t, n, sick, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	sick := []int{1}
	expect := 3
	runSample(t, n, sick, expect)
}
