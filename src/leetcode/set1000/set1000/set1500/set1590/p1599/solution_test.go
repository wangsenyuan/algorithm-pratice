package p1599

import "testing"

func runSample(t *testing.T, cust []int, boarding int, running int, expect int) {
	res := minOperationsMaxProfit(cust, boarding, running)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cust := []int{8, 3}
	boarding := 5
	running := 6
	expect := 3
	runSample(t, cust, boarding, running, expect)
}

func TestSample2(t *testing.T) {
	cust := []int{10, 9, 6}
	boarding := 6
	running := 4
	expect := 7
	runSample(t, cust, boarding, running, expect)
}

func TestSample3(t *testing.T) {
	cust := []int{3, 4, 0, 5, 1}
	boarding := 1
	running := 92
	expect := -1
	runSample(t, cust, boarding, running, expect)
}

func TestSample4(t *testing.T) {
	cust := []int{10, 10, 6, 4, 7}
	boarding := 3
	running := 8
	expect := 9
	runSample(t, cust, boarding, running, expect)
}
func TestSample5(t *testing.T) {
	cust := []int{2}
	boarding := 2
	running := 4
	expect := -1
	runSample(t, cust, boarding, running, expect)
}
