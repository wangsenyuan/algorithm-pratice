package main

import "testing"

func runSample(t *testing.T, x int, y int, expect int) {
	res := solve(x, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 0, 0, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 0, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 1, 0)
}
func TestSample4(t *testing.T) {
	runSample(t, 0, 1, 0)
}

//func TestSample5(t *testing.T) {
//	runSample(t, 3, 3, 1)
//}

//func TestSample6(t *testing.T) {
//	runSample(t, -3, -3, 1)
//}

func TestSample7(t *testing.T) {
	runSample(t, -1, 2, 2)
}

func TestSample8(t *testing.T) {
	runSample(t, 0, -3, 2)
}
