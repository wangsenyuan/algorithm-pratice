package lcp63

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, num int, plate []string, expect [][]int) {
	res := ballGame(num, plate)

	sort.Slice(expect, func(i, j int) bool {
		return expect[i][0] < expect[j][0] || expect[i][0] == expect[j][0] && expect[i][1] < expect[j][1]
	})

	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0] || res[i][0] == res[j][0] && res[i][1] < res[j][1]
	})

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := 4
	plate := []string{
		"..E.", ".EOW", "..W.",
	}
	expect := [][]int{{2, 1}}
	runSample(t, num, plate, expect)
}

func TestSample2(t *testing.T) {
	num := 5
	plate := []string{
		".....", "..E..", ".WO..", ".....",
	}
	expect := [][]int{{0, 1}, {1, 0}, {2, 4}, {3, 2}}
	runSample(t, num, plate, expect)
}

func TestSample3(t *testing.T) {
	num := 3
	plate := []string{
		".....", "....O", "....O", ".....",
	}
	var expect [][]int
	runSample(t, num, plate, expect)
}

func TestSample4(t *testing.T) {
	num := 41
	plate := []string{
		"E...W..WW", ".E...O...", "...WO...W", "..OWW.O..", ".W.WO.W.E", "O..O.W...", ".OO...W..", "..EW.WEE.",
	}
	expect := [][]int{
		{0, 2}, {0, 3}, {0, 5}, {0, 6}, {1, 0}, {1, 8}, {3, 0}, {3, 8}, {4, 0}, {6, 0}, {7, 1}, {7, 4},
	}
	runSample(t, num, plate, expect)
}
