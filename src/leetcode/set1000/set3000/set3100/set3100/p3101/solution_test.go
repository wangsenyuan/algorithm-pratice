package p3101

import "testing"

func runSample(t *testing.T, numBottles, numExchange int, expect int) {
	res := maxBottlesDrunk(numBottles, numExchange)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 13, 6, 15)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 3, 13)
}
