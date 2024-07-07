package p3211

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, expect []string) {
	res := validStrings(n)

	sort.Strings(expect)
	sort.Strings(res)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	expect := []string{"010", "011", "101", "110", "111"}
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 1
	expect := []string{"0", "1"}
	runSample(t, n, expect)
}
