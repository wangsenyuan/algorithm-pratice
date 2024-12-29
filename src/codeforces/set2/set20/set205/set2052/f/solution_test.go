package main

import "testing"

func runSample(t *testing.T, parquet []string, expect string) {
	res := solve(parquet)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []string{
		"#.......##",
		"##..#.##..",
	}
	expect := "Unique"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := []string{
		"...#..",
		"..#...",
	}
	expect := "None"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := []string{
		"........",
		"........",
	}
	expect := "Multiple"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := []string{
		"###",
		"###",
	}
	expect := "Unique"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := []string{
		"..#",
		"..#",
	}
	expect := "Multiple"
	runSample(t, s, expect)
}