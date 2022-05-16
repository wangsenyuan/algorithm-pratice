package p2273

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, words []string, expect []string) {
	res := removeAnagrams(words)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"a", "b", "c", "d", "e"}
	expect := []string{"a", "b", "c", "d", "e"}
	runSample(t, words, expect)
}
