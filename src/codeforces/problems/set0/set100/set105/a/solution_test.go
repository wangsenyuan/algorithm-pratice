package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, k int, before []string, after []string, expect []string) {
	res := solve(k, before, after)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 75
	before := []string{
		"axe 350",
		"impaler 300",
		"ionize 80",
		"megafire 120",
		"magicboost 220",
	}
	after := []string{
		"heal",
		"megafire",
		"shield",
		"magicboost",
	}
	expect := []string{
		"axe 262",
		"heal 0",
		"impaler 225",
		"magicboost 165",
		"megafire 0",
		"shield 0",
	}
	runSample(t, k, before, after, expect)
}

func TestSample2(t *testing.T) {
	k := 94
	before := []string{
		"a 8700",
	}
	after := []string{
		"b",
	}
	expect := []string{
		"a 8178",
		"b 0",
	}
	runSample(t, k, before, after, expect)
}

func TestSample3(t *testing.T) {
	k := 29
	before := []string{
		"a 100",
		"b 200",
		"c 300",
		"d 400",
		"e 500",
		"f 600",
		"g 700",
		"h 800",
		"i 900",
		"j 1000",
		"k 1100",
		"l 1200",
		"m 1300",
		"n 1400",
		"o 1500",
		"p 1600",
		"q 1700",
		"r 1800",
		"s 1900",
		"t 2000",
	}
	after := []string{
		"z",
		"m",
		"k",
	}
	expect := []string{
		"d 116",
		"e 145",
		"f 174",
		"g 203",
		"h 232",
		"i 261",
		"j 290",
		"k 319",
		"l 348",
		"m 377",
		"n 406",
		"o 435",
		"p 464",
		"q 493",
		"r 522",
		"s 551",
		"t 580",
		"z 0",
	}
	runSample(t, k, before, after, expect)
}
