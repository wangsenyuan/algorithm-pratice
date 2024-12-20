package main

import "testing"

func runSample(t *testing.T, num string, expect string) {
	res := solve(num)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}

}

func TestSample1(t *testing.T) {
	runSample(t, "1234567", "1801")
}

func TestSample2(t *testing.T) {
	runSample(t, "101", "11")
}

func TestSample3(t *testing.T) {
	runSample(t, "67378957561978988538", "8716884294")
}

func TestSample4(t *testing.T) {
	runSample(t, "120000", "20001")
}
