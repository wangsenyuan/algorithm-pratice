package p3439

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := maxDifference(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "12233"
	k := 4
	expect := -1
	runSample(t, s, k, expect)
}


func TestSample2(t *testing.T) {
	s := "1122211"
	k := 3
	expect := 1
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "110"
	k := 3
	expect := -1
	runSample(t, s, k, expect)
}
