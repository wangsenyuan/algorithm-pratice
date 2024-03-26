package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, mat []string, expect []string) {
	res := solve(mat)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := []string{
		"*.*",
		".*.",
		"*.*",
	}
	expect := []string{
		"3.3",
		".5.",
		"3.3",
	}
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := []string{
		"**..*",
		"..***",
		".*.*.",
		"*.*.*",
	}
	expect := []string{
		"46..3",
		"..732",
		".6.4.",
		"5.4.3",
	}
	runSample(t, mat, expect)
}
