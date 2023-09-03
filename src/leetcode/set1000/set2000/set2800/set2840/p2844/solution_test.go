package p2844

import (
	"testing"
)

func runSample(t *testing.T, num string, expect int) {
	res := minimumOperations(num)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := "2245047"
	expect := 2
	runSample(t, num, expect)
}

func TestSample2(t *testing.T) {
	num := "25"
	expect := 0
	runSample(t, num, expect)
}

func TestSample3(t *testing.T) {
	num := "75"
	expect := 0
	runSample(t, num, expect)
}
