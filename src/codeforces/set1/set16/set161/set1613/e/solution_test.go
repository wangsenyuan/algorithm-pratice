package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, field []string, expect []string) {
	res := solve(field)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	field := []string{
		"...",
		".L.",
		"...",
	}
	expect := []string{
		"...",
		".L.",
		"...",
	}
	runSample(t, field, expect)
}

func TestSample2(t *testing.T) {
	field := []string{
		"#....",
		"..##L",
		"...#.",
		".....",
	}
	expect := []string{
		"#++++",
		"..##L",
		"...#+",
		"...++",
	}
	runSample(t, field, expect)
}

func TestSample3(t *testing.T) {
	field := []string{
		"....L..#.",
	}
	expect := []string{
		"++++L++#.",
	}
	runSample(t, field, expect)
}
