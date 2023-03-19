package p2591

import "testing"

func runSample(t *testing.T, money int, children int, expect int) {
	res := distMoney(money, children)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	money := 13
	children := 3
	expect := 1
	runSample(t, money, children, expect)
}

func TestSample2(t *testing.T) {
	money := 17
	children := 2
	expect := 1
	runSample(t, money, children, expect)
}

func TestSample3(t *testing.T) {
	money := 23
	children := 2
	expect := 1
	runSample(t, money, children, expect)
}
