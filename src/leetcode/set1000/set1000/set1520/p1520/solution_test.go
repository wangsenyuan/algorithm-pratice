package p1520

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	sort.Strings(expect)
	res := maxNumOfSubstrings(s)

	sort.Strings(res)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %s, expect %v, but bot %v", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abab"
	expect := []string{"abab"}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "adefaddaccc"
	expect := []string{"e", "f", "ccc"}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "abbaccd"
	expect := []string{"d", "bb", "cc"}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "ababa"
	expect := []string{"ababa"}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "zyz"
	expect := []string{"y"}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "cabaaac"
	expect := []string{"b"}
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "cbaabaabc"
	expect := []string{"baabaab"}
	runSample(t, s, expect)
}
