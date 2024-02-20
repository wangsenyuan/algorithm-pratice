package main

import "testing"

func runSample(t *testing.T, s []string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []string{
		"5", "93746", "59", "3746", "593", "746", "5937", "46", "59374", "6",
	}
	expect := 20
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := []string{
		"1", "1",
	}
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := []string{
		"2", "22", "222", "2222", "22222",
	}
	expect := 13
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := []string{
		"12", "21",
	}
	// 1221, 1212, 2121, 2112
	expect := 4
	runSample(t, s, expect)
}
