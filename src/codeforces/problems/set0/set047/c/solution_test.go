package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, words []string, expect []string) {
	res := solve(words)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{
		"NOD",
		"BAA",
		"YARD",
		"AIRWAY",
		"NEWTON",
		"BURN",
	}
	expect := []string{
		"BAA...",
		"U.I...",
		"R.R...",
		"NEWTON",
		"..A..O",
		"..YARD",
	}
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{
		"AAA",
		"AAA",
		"AAAAA",
		"AAA",
		"AAA",
		"AAAAA",
	}
	expect := []string{
		"AAA..",
		"A.A..",
		"AAAAA",
		"..A.A",
		"..AAA",
	}
	runSample(t, words, expect)
}

func TestSample3(t *testing.T) {
	words := []string{
		"PTC",
		"JYNYFDSGI",
		"ZGPPC",
		"IXEJNDOP",
		"JJFS",
		"SSXXQOFGJUZ",
	}
	expect := []string{
		"JJFS....",
		"Y..S....",
		"N..X....",
		"Y..X....",
		"F..Q....",
		"D..O....",
		"S..F....",
		"G..G....",
		"IXEJNDOP",
		"...U...T",
		"...ZGPPC",
	}
	runSample(t, words, expect)
}
