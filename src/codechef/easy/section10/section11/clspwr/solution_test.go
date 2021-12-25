package main

import "testing"

func runSample(t *testing.T, n int64, expect int64) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 9)
}

func TestSample3(t *testing.T) {
	runSample(t, 26, 25)
}

func TestSample4(t *testing.T) {
	runSample(t, 242, 243)
}

func TestSample5(t *testing.T) {
	runSample(t, 129, 128)
}
func TestSample6(t *testing.T) {
	runSample(t, 394857629456789876, 394857628993920400)
}

func TestSample7(t *testing.T) {
	runSample(t, 353872815358409997, 353872815358410000)
}
