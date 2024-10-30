package main

import "testing"

func runSample(t *testing.T, cards []string, expect bool) {
	res := solve(cards)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cards := []string{"2S", "2S", "2C", "2C"}
	expect := true
	runSample(t, cards, expect)
}

func TestSample2(t *testing.T) {
	cards := []string{"2S", "2S", "4S", "3S", "2S"}
	expect := true
	runSample(t, cards, expect)
}

func TestSample3(t *testing.T) {
	cards := []string{"5S", "5S", "7S", "4S", "3H"}
	expect := false
	runSample(t, cards, expect)
}

func TestSample4(t *testing.T) {
	cards := []string{"2S", "4S", "3S", "2S", "4S", "3S", "4S", "4S", "8S", "3S", "2S", "2S", "5S", "3S", "3S", "2S", "3S", "5S", "4S", "4S", "2S", "2S", "4S", "4S", "6S", "2S", "5S", "2S", "5S", "2S", "2S", "2S", "4S", "2S", "5S", "5S", "2S", "6S", "8S", "6S", "2S", "2S", "TS", "2H", "4S", "4S", "3S", "3S", "2S", "2S", "7S", "3S"}
	expect := true
	runSample(t, cards, expect)
}

func TestSample5(t *testing.T) {
	cards := []string{"JS", "7S", "3S", "2S", "7S", "TS", "4S", "6S", "5H", "TS", "4S", "TH", "6H", "9S", "TH", "TH", "4S", "4H", "2H", "TH", "TC", "TH", "TS", "TS", "4S", "TS", "2S", "TH", "TH", "TS", "6S", "TS", "TS", "3S", "TS", "TH", "5H", "TS", "TS", "5S", "7H", "2H", "TS", "6S", "6H", "2H", "TS", "TH", "2S", "4S", "4H", "4S"}
	expect := true
	runSample(t, cards, expect)
}
