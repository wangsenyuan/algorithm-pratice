package p5563

import "testing"

func runSample(t *testing.T, inventory []int, orders int, expect int) {
	res := maxProfit(inventory, orders)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	inventory := []int{2, 5}
	orders := 4
	expect := 14
	runSample(t, inventory, orders, expect)
}

func TestSample2(t *testing.T) {
	inventory := []int{3, 5}
	orders := 6
	expect := 19
	runSample(t, inventory, orders, expect)
}

func TestSample3(t *testing.T) {
	inventory := []int{2, 8, 4, 10, 6}
	orders := 20
	expect := 110
	runSample(t, inventory, orders, expect)
}

func TestSample4(t *testing.T) {
	inventory := []int{1000000000}
	orders := 1000000000
	expect := 21
	runSample(t, inventory, orders, expect)
}

