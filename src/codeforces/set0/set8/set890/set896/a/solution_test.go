package main

import (
	"bytes"
	"testing"
)

func runSample(t *testing.T, queries [][]int, expect string) {
	var buf bytes.Buffer
	for _, cur := range queries {
		n, k := cur[0], cur[1]
		buf.WriteString(solve(n, k))
	}

	res := buf.String()

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	queries := [][]int{
		{1, 1},
		{1, 2},
		{1, 111111111111},
	}
	expect := "Wh."

	runSample(t, queries, expect)
}

func TestSample2(t *testing.T) {
	queries := [][]int{
		{0, 69},
		{1, 194},
		{1, 139},
		{0, 47},
		{1, 66},
	}
	expect := "abdef"

	runSample(t, queries, expect)
}

func TestSample3(t *testing.T) {
	queries := [][]int{
		{4, 1825},
		{3, 75},
		{3, 530},
		{4, 1829},
		{4, 1651},
		{3, 187},
		{4, 584},
		{4, 255},
		{4, 774},
		{2, 474},
	}
	expect := "Areyoubusy"

	runSample(t, queries, expect)
}

func TestSample4(t *testing.T) {
	queries := [][]int{
		{999, 1000000000000000000},
	}
	expect := "?"

	runSample(t, queries, expect)
}
