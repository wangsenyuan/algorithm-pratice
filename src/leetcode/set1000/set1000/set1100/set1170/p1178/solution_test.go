package p1178

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, words []string, puzzles []string, expect []int) {
	res := findNumOfValidWords(words, puzzles)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", words, puzzles, expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"}
	puzzles := []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}
	expect := []int{1, 1, 3, 2, 4, 0}
	runSample(t, words, puzzles, expect)
}
