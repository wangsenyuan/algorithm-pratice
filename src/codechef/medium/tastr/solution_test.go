package main

import (
	"testing"
	"reflect"
)

func runUniqueSubStrCount(t *testing.T, s string, expect int64) {
	res := uniqueSubStrCount([]byte(s))
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestUniqueSubStrCount1(t *testing.T) {
	runUniqueSubStrCount(t, "abaabba", 21)
}

func runSample(t *testing.T, a, b string, expect int64) {
	res := solve([]byte(a), []byte(b))

	if res != expect {
		t.Errorf("Sample %s %s, expect %d, but got %d", a, b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aacd", "cdaa", 8)
}

func runBuildSuffixArray(t *testing.T, s string, expect []int) {
	res, _, _ := buildSuffixArray(len(s), []byte(s))
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %s, expect %v, but got %v", s, expect, res)
	}
}

func TestBuilidSuffixArray1(t *testing.T) {
	runBuildSuffixArray(t, "banana", []int{5, 3, 1, 0, 4, 2})
}
