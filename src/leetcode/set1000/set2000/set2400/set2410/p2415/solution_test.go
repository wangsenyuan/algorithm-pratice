package p2415

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, words []string, expect []int) {
	res := sumPrefixScores(words)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"abc", "ab", "bc", "b"}
	expect := []int{5, 4, 3, 2}
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{"abcd"}
	expect := []int{4}
	runSample(t, words, expect)
}
