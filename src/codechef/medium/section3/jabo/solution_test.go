package main

import "testing"

func runSample(t *testing.T, n, r, c int, steps []string, expect []bool) {
	solver := NewSolver(r, c)

	var j int

	for i := 0; i < n; i++ {
		step := steps[i]
		if step[0] == 'W' {
			solver.ConnectWire(step[1:])
		} else if step[0] == 'V' {
			solver.AddVoltage(step[1:])
		} else if step[0] == 'R' {
			solver.RemoveVoltage(step[1:])
		} else {
			ans := solver.CheckLed(step[1:])
			if ans != expect[j] {
				t.Fatalf("sample %d %d %d %v, expect %v, but fail at %d", n, r, c, steps, expect, j)
			}
			j++
		}
	}
}

func TestSample1(t *testing.T) {
	n, r, c := 9, 2, 10
	steps := []string{
		"WADAEAFAG",
		"VAGAD",
		"LAFAHAGAE",
		"VAKAK",
		"LAJAKAKAJ",
		"LAKAIAGAB",
		"LABABACAB",
		"RAKAK",
		"LAJAKAKAJ",
	}
	expect := []bool{true, true, false, false, false}
	runSample(t, n, r, c, steps, expect)
}
