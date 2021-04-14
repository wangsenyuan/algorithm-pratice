package lcp37

import (
	"math"
	"testing"
)

func runSample(t *testing.T, lines [][]int, expect float64) {
	res := minRecSize(lines)

	if math.Abs(res-expect) > eps {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	lines := [][]int{
		{2, 3}, {3, 0}, {4, 1},
	}
	var expect float64 = 48
	runSample(t, lines, expect)
}
