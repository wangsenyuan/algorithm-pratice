package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, G []string, queries []int, expect []string) {
	res := solve(n, G, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v, but got %v", n, G, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	G := []string{
		"2+1-2",
		"+3-4+",
		"5+2+1",
		"-4-0-",
		"9+5+1",
	}
	queries := []int{20, 30, 40}
	expect := []string{
		"1+5+5+9",
		"3+4+5+9+9",
		"4+9+9+9+9",
	}
	runSample(t, n, G, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	G := []string{
		"2+1",
		"+4+",
		"5+1",
	}
	queries := []int{2, 20}
	expect := []string{
		"2", "5+5+5+5",
	}
	runSample(t, n, G, queries, expect)
}
