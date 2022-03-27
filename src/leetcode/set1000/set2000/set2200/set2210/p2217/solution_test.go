package p2217

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, queries []int, intLength int, expect []int64) {
	res := kthPalindrome(queries, intLength)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	qs := []int{1, 2, 3, 4, 5, 90}
	l := 3
	expect := []int64{101, 111, 121, 131, 141, 999}
	runSample(t, qs, l, expect)
}

func TestSample2(t *testing.T) {
	qs := []int{2, 4, 6}
	l := 4
	expect := []int64{1111, 1331, 1551}
	runSample(t, qs, l, expect)
}

func TestSample3(t *testing.T) {
	qs := []int{1, 2, 3, 4, 5, 90}
	l := 15
	expect := []int64{100000000000001, 100000010000001, 100000020000001, 100000030000001, 100000040000001, 100000898000001}
	runSample(t, qs, l, expect)
}

func TestSample4(t *testing.T) {
	qs := []int{1, 2, 3, 4, 5, 90}
	l := 5
	expect := []int64{10001, 10101, 10201, 10301, 10401, 18981}
	runSample(t, qs, l, expect)
}

func TestSample5(t *testing.T) {
	qs := []int{1, 2, 3, 4, 5, 90}
	l := 6
	expect := []int64{100001, 101101, 102201, 103301, 104401, 189981}
	runSample(t, qs, l, expect)
}

func TestSample6(t *testing.T) {
	qs := []int{2, 201429812, 8, 520498110, 492711727, 339882032, 462074369, 9, 7, 6}
	l := 1
	expect := []int64{2, -1, 8, -1, -1, -1, -1, 9, 7, 6}
	runSample(t, qs, l, expect)
}

func TestSample7(t *testing.T) {
	qs := []int{2, 201429812, 8, 520498110, 492711727, 339882032, 462074369, 9, 7, 6}
	l := 2
	expect := []int64{22, -1, 88, -1, -1, -1, -1, 99, 77, 66}
	runSample(t, qs, l, expect)
}

func TestSample8(t *testing.T) {
	qs := []int{98043237}
	l := 15
	expect := []int64{-1}
	runSample(t, qs, l, expect)
}
