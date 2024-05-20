package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, trump string, cards []string, expect bool) {
	res := solve(trump, cards)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	check := func(a string, b string) bool {
		// can cards[j] beats i
		if a[1] == b[1] {
			return a[0] < b[0]
		}
		return b[1] == trump[0]
	}

	for i := 0; i < len(res); i += 2 {
		if !check(res[i], res[i+1]) {
			t.Fatalf("Sample result %v not correct", res)
		}
	}

	sort.Strings(res)
	sort.Strings(cards)

	if !reflect.DeepEqual(res, cards) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

}

func TestSample1(t *testing.T) {
	trump := "S"
	cards := []string{
		"3C", "9S", "4C", "6D", "3S", "7S",
	}
	expect := true
	runSample(t, trump, cards, expect)
}

func TestSample2(t *testing.T) {
	trump := "C"
	cards := []string{
		"3S", "5D", "9S", "6H",
	}
	expect := false
	runSample(t, trump, cards, expect)
}

func TestSample3(t *testing.T) {
	trump := "H"
	cards := []string{
		"6C", "5D",
	}
	expect := false
	runSample(t, trump, cards, expect)
}

func TestSample4(t *testing.T) {
	trump := "S"
	cards := []string{
		"7S", "3S",
	}
	expect := true
	runSample(t, trump, cards, expect)
}

func TestSample5(t *testing.T) {
	trump := "H"
	cards := []string{
		"9S", "9H",
	}
	expect := true
	runSample(t, trump, cards, expect)
}

func TestSample6(t *testing.T) {
	trump := "H"
	cards := []string{
		"9S", "9H",
	}
	expect := true
	runSample(t, trump, cards, expect)
}

func TestSample7(t *testing.T) {
	trump := "C"
	cards := []string{
		"9C", "9S", "6H", "8C",
	}
	expect := true
	runSample(t, trump, cards, expect)
}

func TestSample8(t *testing.T) {
	trump := "H"
	cards := []string{
		"8D", "2H", "2C", "9H",
	}
	expect := true
	runSample(t, trump, cards, expect)
}

func TestSample9(t *testing.T) {
	trump := "S"
	cards := []string{
		"5S", "9D", "6D", "9C",
	}
	expect := true
	runSample(t, trump, cards, expect)
}
