package p1801

import "testing"

func runSample(t *testing.T, orders [][]int, expect int) {
	res := getNumberOfBacklogOrders(orders)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	orders := [][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}}
	expect := 6
	runSample(t, orders, expect)
}

func TestSample2(t *testing.T) {
	orders := [][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}}
	expect := 999999984
	runSample(t, orders, expect)
}
