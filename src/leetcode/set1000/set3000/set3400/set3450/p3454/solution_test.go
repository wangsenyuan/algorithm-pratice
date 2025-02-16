package p3454

import (
	"math"
	"testing"
)

func runSample(t *testing.T, squares [][]int, expect float64) {
	res := separateSquares(squares)

	if math.Abs(res-expect) > 1e-5 {
		t.Fatalf("Sample expect %.7f, but got %.7f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	squares := [][]int{
		{0, 0, 1}, {2, 2, 1},
	}
	expect := 1.0
	runSample(t, squares, expect)
}

func TestSample2(t *testing.T) {
	squares := [][]int{
		{0, 0, 2}, {1, 1, 1},
	}
	expect := 1.0
	runSample(t, squares, expect)
}

func TestSample3(t *testing.T) {
	squares := [][]int{
		{15, 21, 2}, {19, 21, 3},
	}
	expect := 22.3
	runSample(t, squares, expect)
}

func TestSample4(t *testing.T) {
	squares := [][]int{
		{26, 29, 3}, {10, 24, 1},
	}
	expect := 30.33333
	runSample(t, squares, expect)
}

func TestSample5(t *testing.T) {
	squares := [][]int{
		{999892931, 999974790, 434471746}, {319710671, 963660807, 875442433}, {353202089, 976938743, 622045959}, {765760000, 939956921, 271907109}, {234214719, 848813522, 26688635}, {154771654, 645515409, 804966565}, {599682863, 948151649, 645372386}, {534582408, 880488308, 338822523}, {211765761, 379959124, 514093466}, {331072193, 790928668, 912992282},
	}
	expect := 1173798088.14062

	runSample(t, squares, expect)
}
