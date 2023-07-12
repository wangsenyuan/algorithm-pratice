package main

import "testing"

func runSample(t *testing.T, price string, expect int) {
	res := solve(price)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	price := "107"
	expect := 42
	runSample(t, price, expect)
}

func TestSample2(t *testing.T) {
	price := "100500100500"
	expect := 428101984
	runSample(t, price, expect)
}
