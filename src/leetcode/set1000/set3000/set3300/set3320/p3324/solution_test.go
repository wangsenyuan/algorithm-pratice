package p3324

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, target string, expect []string) {
	res := stringSequence(target)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	target := "abc"
	expect := []string{"a", "aa", "ab", "aba", "abb", "abc"}
	runSample(t, target, expect)
}
