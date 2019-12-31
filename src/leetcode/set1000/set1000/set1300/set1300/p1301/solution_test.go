package p1301

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, board []string, expect []int) {
	res := pathsWithMaxScore(board)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", board, expect, res)
	}
}

func TestSample1(t *testing.T) {
	board := []string{
		"E23", "2X2", "12S",
	}
	expect := []int{7, 1}
	runSample(t, board, expect)
}

func TestSample2(t *testing.T) {
	board := []string{
		"E12", "1X1", "21S",
	}
	expect := []int{4, 2}
	runSample(t, board, expect)
}

func TestSample3(t *testing.T) {
	board := []string{
		"E11", "XXX", "11S",
	}
	expect := []int{0, 0}
	runSample(t, board, expect)
}
