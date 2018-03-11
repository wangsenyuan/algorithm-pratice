package p799

import (
	"testing"
	"math"
)

func runSample(t *testing.T, poured int, query_row int, query_glass int, expect float64) {
	res := champagneTower(poured, query_row, query_glass)

	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("sample %d %d %d, expect %.7f, but got %.7f", poured, query_row, query_glass, expect, res)
	}
}

func TestSample1(t *testing.T) {
	poured := 1
	query_glass := 1
	query_row := 1
	expect := 0.0
	runSample(t, poured, query_row, query_glass, expect)
}

func TestSample2(t *testing.T) {
	poured := 2
	query_glass := 1
	query_row := 1
	expect := 0.5
	runSample(t, poured, query_row, query_glass, expect)
}
