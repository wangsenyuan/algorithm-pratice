package main

import (
	"math"
	"testing"
)

func runSample(tc *testing.T, f, R, t, r, g float64, expect float64) {
	res := solve(f, R, t, r, g)
	if math.Abs(res-expect) > 1e-6 {
		tc.Errorf("Sample %f %f %f %f %f, expect %.8f, but got %.8f", f, R, t, r, g, expect, res)
	}
}

func TestSample1(tc *testing.T) {
	f, R, t, r, g := 0.25, 1.0, 0.1, 0.01, 0.5
	expect := 1.0
	runSample(tc, f, R, t, r, g, expect)
}

func TestSample2(tc *testing.T) {
	f, R, t, r, g := 0.25, 1.0, 0.1, 0.01, 0.9
	expect := 0.910015
	runSample(tc, f, R, t, r, g, expect)
}

func TestSample3(tc *testing.T) {
	f, R, t, r, g := 0.00001, 10000.0, 0.00001, 0.00001, 1000.0
	expect := 0.000000
	runSample(tc, f, R, t, r, g, expect)
}

func TestSample4(tc *testing.T) {
	f, R, t, r, g := 0.4, 10000.0, 0.00001, 0.00001, 700.0
	expect := 0.002371
	runSample(tc, f, R, t, r, g, expect)
}

func TestSample5(tc *testing.T) {
	f, R, t, r, g := 1.0, 100.0, 1.0, 1.0, 10.0
	expect := 0.573972
	runSample(tc, f, R, t, r, g, expect)
}
