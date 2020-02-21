package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, G []string, expect []string) {
	res := solve(n, G)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	G := []string{
		"01",
		"01",
	}

	expect := []string{
		"10",
		"01",
	}

	runSample(t, n, G, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	G := []string{
		"0111",
		"0001",
		"1010",
		"1010",
	}

	expect := []string{
		"1110",
		"0001",
		"1010",
		"1010",
	}

	runSample(t, n, G, expect)
}
