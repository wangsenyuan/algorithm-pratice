package p1467

import (
	"math"
	"testing"
)

func runSample(t *testing.T, balls []int, expect float64) {
	res := getProbability(balls)

	if math.Abs(res-expect) > 1e-5 {
		t.Errorf("Sample %v, expect %f, but got %f", balls, expect, res)
	}
}

func TestSample1(t *testing.T) {
	balls := []int{1, 1}
	expect := 1.0
	runSample(t, balls, expect)
}

func TestSample2(t *testing.T) {
	balls := []int{2, 1, 1}
	expect := 0.66667
	runSample(t, balls, expect)
}

func TestSample3(t *testing.T) {
	balls := []int{1, 2, 1, 2}
	expect := 0.6
	runSample(t, balls, expect)
}

func TestSample4(t *testing.T) {
	balls := []int{3, 2, 1}
	expect := 0.3
	runSample(t, balls, expect)
}

func TestSample5(t *testing.T) {
	balls := []int{6, 6, 6, 6, 6, 6}
	expect := 0.90327
	runSample(t, balls, expect)
}
