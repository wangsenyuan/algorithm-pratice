package p950

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, deck []int, expect []int) {
	res := deckRevealedIncreasing(deck)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", deck, expect, res)
	}
}

func TestSample1(t *testing.T) {
	deck := []int{17, 13, 11, 2, 3, 5, 7}
	expect := []int{2, 13, 3, 11, 5, 17, 7}
	runSample(t, deck, expect)
}
