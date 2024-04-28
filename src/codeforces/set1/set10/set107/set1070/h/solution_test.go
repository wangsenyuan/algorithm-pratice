package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, f []string, q []string, expect []int) {
	cnt, _ := solve(f, q)

	if !reflect.DeepEqual(expect, cnt) {
		t.Fatalf("Sample expect %v, but got %v", expect, cnt)
	}
}

func TestSample1(t *testing.T) {
	f := []string{
		"test",
		"contests",
		"test.",
		".test",
	}
	q := []string{
		"ts",
		".",
		"st.",
		".test",
		"contes.",
		"st",
	}
	expect := []int{1, 2, 1, 1, 0, 4}
	runSample(t, f, q, expect)
}

func TestSample2(t *testing.T) {
	f := []string{
		"0",
		"00",
		"000",
		"0000",
		"00000",
		"000000",
		"0000000",
		"00000000",
	}
	q := []string{
		"0",
		"00",
		"000",
		"0000",
		"00000",
		"000000",
		"0000000",
		"00000000",
	}
	expect := []int{8, 7, 6, 5, 4, 3, 2, 1}
	runSample(t, f, q, expect)
}
