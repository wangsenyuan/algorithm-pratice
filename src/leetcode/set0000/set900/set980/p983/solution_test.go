package p983

import "testing"

func runSample(t *testing.T, days []int, costs []int, expect int) {
	res := mincostTickets(days, costs)

	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", days, costs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}, 11)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}, 17)
}
