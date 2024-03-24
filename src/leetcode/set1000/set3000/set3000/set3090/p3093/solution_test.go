package p3093

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, words []string, query []string, expect []int) {
	res := stringIndices(words, query)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"abcd", "bcd", "xbcd"}
	query := []string{"cd", "bcd", "xyz"}
	expect := []int{1, 1, 1}
	runSample(t, words, query, expect)
}

func TestSample2(t *testing.T) {
	words := []string{"abcdefgh", "poiuygh", "ghghgh"}
	query := []string{"gh", "acbfgh", "acbfegh"}
	expect := []int{2, 0, 2}
	runSample(t, words, query, expect)
}
