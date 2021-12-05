package p2093

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, arr []int, expect []int) {
	res := findEvenNumbers(arr)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{2, 2, 8, 8, 2}
	expect := []int{222, 228, 282, 288, 822, 828, 882}
	runSample(t, arr, expect)
}
