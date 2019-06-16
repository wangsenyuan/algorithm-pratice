package p289

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	board := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}
	expect := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	gameOfLife(board)

	if !reflect.DeepEqual(board, expect) {
		t.Errorf("expect %v, but got %v", expect, board)
	}
}
