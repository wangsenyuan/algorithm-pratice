package main

import "testing"

func runSample(t *testing.T, str string, expect int64) {
	res := solve(str)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", str, expect, res)
	}
}

func TestSample1(t *testing.T) {
	str := "##.#.."
	runSample(t, str, 1)
}

func TestSample2(t *testing.T) {
	str := ".#.#.#."
	runSample(t, str, 6)
}

func TestSample3(t *testing.T) {
	str := "##.##"
	runSample(t, str, -1)
}
