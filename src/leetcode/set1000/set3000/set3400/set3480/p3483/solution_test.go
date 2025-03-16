package p3483

import "testing"

func runSample(t *testing.T, digits []int, expect int) {
	res := totalNumbers(digits)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	digits := []int{4, 7, 7, 3}
	expect := 3
	runSample(t, digits, expect)
}
