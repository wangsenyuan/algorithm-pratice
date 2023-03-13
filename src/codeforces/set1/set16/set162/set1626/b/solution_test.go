package main

import (
	"fmt"
	"testing"
)

func runSample(t *testing.T, x string, expect string) {
	res := solve(x)

	if res != expect {
		t.Fatalf("Sample  %s expect %s, but got %s", x, expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := "10057"
	expect := "10012"
	runSample(t, x, expect)
}

func TestSample2(t *testing.T) {
	x := "90"
	expect := "9"
	runSample(t, x, expect)
}

func TestSample3(t *testing.T) {
	x := "100"
	expect := "10"
	runSample(t, x, expect)
}

func TestSample4(t *testing.T) {
	x := "191"
	expect := "110"
	runSample(t, x, expect)
}

func TestSample6(t *testing.T) {
	x := "190"
	expect := "100"
	runSample(t, x, expect)
}

func TestSample5(t *testing.T) {
	for x := 100; x < 10000; x++ {
		s := fmt.Sprintf("%d", x)
		expect := bruteForce(s)
		runSample(t, s, expect)
	}
}

func TestSample7(t *testing.T) {
	x := "2819"
	expect := "2810"
	runSample(t, x, expect)
}
