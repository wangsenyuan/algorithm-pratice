package p3352

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := countKReducibleNumbers(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "111"
	k := 1
	expect := 3
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "1000"
	k := 2
	expect := 6
	runSample(t, s, k, expect)
}
