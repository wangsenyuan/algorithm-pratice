package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, P, S []string, expect []int) {
	res := solve(P, S)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []string{
		"he",
		"she",
		"sher",
		"his",
		"hers",
	}
	S := []string{
		"ushers",
		"she-said-he-said-she-said-he-said-his",
	}
	expect := []int{5, 3, 1, 1, 1}
	runSample(t, P, S, expect)
}

func TestSample2(t *testing.T) {
	P := []string{
		"who",
		"shawty",
		"hawt",
	}
	S := []string{
		"Get-it-shawty-Get-it-shawty",
		"Whoa-W-W-Whoa-Shawtyyyyy",
	}
	expect := []int{0, 2, 3}
	runSample(t, P, S, expect)
}
