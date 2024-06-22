package main

import "testing"

func runSample(t *testing.T, grid []string, expect_string string, expect_cnt int) {
	res, cnt := solve(grid)

	if res != expect_string || cnt != expect_cnt {
		t.Fatalf("Sample expect %s, %d, but got %s, %d", expect_string, expect_cnt, res, cnt)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"00",
		"00",
	}
	expect_string := "000"
	expect_cnt := 2
	runSample(t, grid, expect_string, expect_cnt)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"1101",
		"1100",
	}
	expect_string := "11000"
	expect_cnt := 1
	runSample(t, grid, expect_string, expect_cnt)
}

func TestSample3(t *testing.T) {
	grid := []string{
		"00100111",
		"11101101",
	}
	expect_string := "001001101"
	expect_cnt := 4
	runSample(t, grid, expect_string, expect_cnt)
}

func TestSample4(t *testing.T) {
	grid := []string{
		"010",
		"010",
	}
	expect_string := "0010"
	expect_cnt := 1
	runSample(t, grid, expect_string, expect_cnt)
}
