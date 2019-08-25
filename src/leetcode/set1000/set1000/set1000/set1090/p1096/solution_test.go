package p1096

import (
	"reflect"
	"sort"
	"testing"
)

func runTestSample(t *testing.T, expr string, expect []string) {
	res := braceExpansionII(expr)

	sort.Strings(res)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %s, expect %v, but got %v", expr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	expr := "{a,b}{c{d,e}}"
	expect := []string{"acd", "ace", "bcd", "bce"}
	runTestSample(t, expr, expect)
}

func TestSample2(t *testing.T) {
	expr := "{{a,z},a{b,c},{ab,z}}"
	expect := []string{"a", "ab", "ac", "z"}
	runTestSample(t, expr, expect)
}

func TestSample3(t *testing.T) {
	expr := "{a{x,ia,o}w,{n,{g{u,o}},{a,{x,ia,o},w}},er}"
	expect := []string{"a", "aiaw", "aow", "axw", "er", "go", "gu", "ia", "n", "o", "w", "x"}
	runTestSample(t, expr, expect)
}
