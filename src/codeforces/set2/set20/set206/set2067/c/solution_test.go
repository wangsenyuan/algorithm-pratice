package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 51, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 60, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 61, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 777, 0)
}

func TestSample5(t *testing.T) {
	runSample(t, 12345689, 1)
}

func TestSample6(t *testing.T) {
	runSample(t, 1000000000, 3)
}

func TestSample7(t *testing.T) {
	runSample(t, 2002, 5)
}

func TestSample8(t *testing.T) {
	runSample(t, 3001, 4)
}

func TestSample9(t *testing.T) {
	runSample(t, 977, 0)
}

func TestSample10(t *testing.T) {
	runSample(t, 989898986, 7)
}

func TestSample11(t *testing.T) {
	runSample(t, 80, 1)
}
func TestSample12(t *testing.T) {
	runSample(t, 800001, 2)
}
func TestSample13(t *testing.T) {
	runSample(t, 96, 7)
}

func TestSample14(t *testing.T) {
	runSample(t, 15, 7)
}

func TestSample15(t *testing.T) {
	runSample(t, 90, 3)
}