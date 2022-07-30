package main

import "testing"

func runSample(t *testing.T, A, B int64, expect int64) {
	res := solve(A, B)
	if res != expect {
		t.Fatalf("Sample %d %d expect %d, but got %d", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 101, 101, 1)
}

func TestSample7(t *testing.T) {
	runSample(t, 100, 102, 2)
}

func TestSample0(t *testing.T) {
	var cnt int64
	for num := int64(100); num <= 200; num++ {
		cnt += count(num)
		runSample(t, 1, num, cnt)
	}
}

func TestSample2(t *testing.T) {
	runSample(t, 100, 150, 19)
}

func TestSample3(t *testing.T) {
	runSample(t, 100, 999, 525)
}

func TestSample5(t *testing.T) {
	runSample(t, 100, 1011, 528)
}

func TestSample6(t *testing.T) {
	runSample(t, 100, 1000, 525)
}
