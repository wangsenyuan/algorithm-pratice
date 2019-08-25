package p1170

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, queries []string, words []string, expect []int) {
	res := numSmallerByFrequency(queries, words)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", queries, words, expect, res)
	}
}

func TestSample1(t *testing.T) {
	queries := []string{"cbd"}
	words := []string{"zaaaz"}
	expect := []int{1}
	runSample(t, queries, words, expect)
}

func TestSample2(t *testing.T) {
	queries := []string{"bbb", "cc"}
	words := []string{"a", "aa", "aaa", "aaaa"}
	expect := []int{1, 2}
	runSample(t, queries, words, expect)
}
