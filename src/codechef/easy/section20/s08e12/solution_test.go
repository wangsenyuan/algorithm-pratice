package main

import "testing"

func runSample(t *testing.T, S, T string, expect int) {
	res := solve(len(S), len(T), S, T)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "1000110001"
	T := "010001100010010001100010"
	expect := 9
	runSample(t, S, T, expect)
}

func TestSample2(t *testing.T) {
	S := "101"
	T := "000"
	expect := 2
	runSample(t, S, T, expect)
}

func TestSample3(t *testing.T) {
	S := "1101"
	T := "100101"
	expect := 3
	runSample(t, S, T, expect)
}

func TestSample4(t *testing.T) {
	S := "010011"
	T := "01010010"
	expect := 4
	runSample(t, S, T, expect)
}
