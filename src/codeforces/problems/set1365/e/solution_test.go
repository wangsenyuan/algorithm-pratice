package main

import "testing"

func runSample(t *testing.T, n int, A []uint64, expect uint64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []uint64{1, 2, 3}
	var expect uint64 = 3
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []uint64{3, 1, 4}
	var expect uint64 = 7
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 1
	A := []uint64{1}
	var expect uint64 = 1
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	A := []uint64{7, 7, 1, 1}
	var expect uint64 = 7
	runSample(t, n, A, expect)
}

func TestSample5(t *testing.T) {
	n := 2
	A := []uint64{3, 4}
	var expect uint64 = 7
	runSample(t, n, A, expect)
}

func TestSample6(t *testing.T) {
	n := 10
	A := []uint64{582366931603099761, 314858607473442114, 530263190370309150, 871012489649491233, 877068367969362781, 671646356752418008, 390155369686708364, 958695211216189893, 919124874293325142, 196726357117434998}
	var expect uint64 = 1152921229728939135
	runSample(t, n, A, expect)
}
