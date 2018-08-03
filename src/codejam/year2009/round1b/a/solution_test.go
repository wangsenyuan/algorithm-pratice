package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, S string, animals [][]string, expect []float64) {
	res := solve(S, animals)
	if len(expect) != len(res) {
		t.Errorf("Sample %s %v, expect %v, but got %v", S, animals, expect, res)
		return
	}
	for i := 0; i < len(expect); i++ {
		if math.Abs(res[i]-expect[i]) > 1e-6 {
			t.Errorf("Sample %s %v, expect %v, but got %v, differ at pos %d", S, animals, expect, res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	S := `
	(0.5 cool
		( 1.000)
		(0.5 ))
	`
	animals := [][]string{
		{"cool"}, {},
	}
	expect := []float64{0.5, 0.25}
	runSample(t, S, animals, expect)
}

func TestSample2(t *testing.T) {
	S := `
	(0.2 furry
		(0.81 fast
			(0.3)
			(0.2)
		)
		(0.1 fishy
			(0.3 freshwater
				(0.01)
				(0.01)
			)
			(0.1)
		)
	)
	`
	animals := [][]string{
		{"furry", "freshwater"}, {"fast", "freshwater", "fishy", "rainbowy"}, {"extinct"},
	}
	expect := []float64{0.0324000, 0.0000600, 0.0020000}
	runSample(t, S, animals, expect)
}
