package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, records []string, expect []string) {
	res := solve(records)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	records := []string{
		"manutd 8 vs. 2 arsenal",
		"lyon 1 vs. 2 manutd",
		"fcbarca 0 vs. 0 lyon",
		"fcbarca 5 vs. 1 arsenal",
		"manutd 3 vs. 1 fcbarca",
		"arsenal 6 vs. 0 lyon",
		"arsenal 0 vs. 0 manutd",
		"manutd 4 vs. 2 lyon",
		"arsenal 2 vs. 2 fcbarca",
		"lyon 0 vs. 3 fcbarca",
		"lyon 1 vs. 0 arsenal",
		"fcbarca 0 vs. 1 manutd",
	}
	expect := []string{"manutd", "fcbarca"}

	runSample(t, records, expect)
}

func TestSample2(t *testing.T) {
	records := []string{
		"a 3 vs. 0 b",
		"a 0 vs. 0 c",
		"a 0 vs. 0 d",
		"b 0 vs. 0 a",
		"b 4 vs. 0 c",
		"b 0 vs. 0 d",
		"c 0 vs. 0 a",
		"c 0 vs. 0 b",
		"c 1 vs. 0 d",
		"d 3 vs. 0 a",
		"d 0 vs. 0 b",
		"d 0 vs. 0 c",
	}
	expect := []string{"d", "b"}

	runSample(t, records, expect)
}
