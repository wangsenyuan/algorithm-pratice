package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, board []string, expect []string) {
	res := solve(board)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	board := []string{
		".8....2..",
		".1....5..",
		"..34..7..",
		"..9.5....",
		".2...46..",
		"3........",
		"9...2....",
		".........",
		"......4.7",
	}
	expect := []string{
		"486715293",
		"712938546",
		"593462718",
		"679251384",
		"128394675",
		"354876129",
		"945627831",
		"867143952",
		"231589467",
	}
	runSample(t, board, expect)
}

func TestSample2(t *testing.T) {
	board := []string{
		"....41...",
		"...6....5",
		".....7.9.",
		"....1.3..",
		".5......1",
		".2.......",
		"..18...76",
		".7......2",
		"........3",
	}
	expect := []string{
		"293541768",
		"748692135",
		"615387294",
		"864715329",
		"357269481",
		"129438657",
		"531824976",
		"476953812",
		"982176543",
	}
	runSample(t, board, expect)
}
