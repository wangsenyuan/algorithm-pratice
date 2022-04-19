package p2240

import "testing"

func runSample(t *testing.T, total int, cost1 int, cost2 int, expect int64) {
	res := waysToBuyPensPencils(total, cost1, cost2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	total := 20
	cost1 := 10
	cost2 := 5
	var expect int64 = 9
	runSample(t, total, cost1, cost2, expect)
}

func TestSample2(t *testing.T) {
	total := 5
	cost1 := 10
	cost2 := 10
	var expect int64 = 1
	runSample(t, total, cost1, cost2, expect)
}
