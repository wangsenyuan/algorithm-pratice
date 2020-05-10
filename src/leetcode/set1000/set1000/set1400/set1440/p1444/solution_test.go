package p1444

import "testing"

func runSample(t *testing.T, pizza []string, k int, expect int) {
	res := ways(pizza, k)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", pizza, expect, res)
	}
}

func TestSample1(t *testing.T) {
	pizza := []string{
		"A..", "AAA", "...",
	}
	k := 3
	expect := 3
	runSample(t, pizza, k, expect)
}

func TestSample2(t *testing.T) {
	pizza := []string{
		"A..", "AA.", "...",
	}
	k := 3
	expect := 1
	runSample(t, pizza, k, expect)
}

func TestSample3(t *testing.T) {
	pizza := []string{
		"A..", "A..", "...",
	}
	k := 1
	expect := 1
	runSample(t, pizza, k, expect)
}
