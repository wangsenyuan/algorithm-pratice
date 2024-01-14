package p3007

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, a string, b string, k int, expect []int) {
	res := beautifulIndices(s, a, b, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "isawsquirrelnearmysquirrelhouseohmy"
	a := "my"
	b := "squirrel"
	k := 15
	expect := []int{16, 33}
	runSample(t, s, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	s := "abcd"
	a := "a"
	b := "a"
	k := 4
	expect := []int{0}
	runSample(t, s, a, b, k, expect)
}
