package p3453

import (
	"math"
	"testing"
)

func runSample(t *testing.T, squares [][]int, expect float64) {
	res := separateSquares(squares)

	if math.Abs(res-expect) > 1e-6 {
		t.Fatalf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	squares := [][]int{{522261215, 954313664, 461744743}, {628661372, 718610752, 21844764}, {619734768, 941310679, 91724451}, {352367502, 656774918, 591943726}, {860247066, 905800565, 853111524}, {817098516, 868361139, 817623995}, {580894327, 654069233, 691552059}, {182377086, 256660052, 911357}, {151104008, 908768329, 890809906}, {983970552, 992192635, 462847045}}
	expect := 1205916532.61256
	runSample(t, squares, expect)
}
