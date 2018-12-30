package p967

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, N, K int, expect []int) {
	res := numsSameConsecDiff(N, K)
	sort.Ints(res)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d, expect %v, but got %v", N, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K := 3, 7
	expect := []int{181, 292, 707, 818, 929}
	sort.Ints(expect)
	runSample(t, N, K, expect)
}

func TestSample2(t *testing.T) {
	N, K := 2, 1
	expect := []int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98}
	sort.Ints(expect)
	runSample(t, N, K, expect)
}
func TestSample3(t *testing.T) {
	N, K := 2, 0
	expect := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	sort.Ints(expect)
	runSample(t, N, K, expect)
}

func TestSample4(t *testing.T) {
	N, K := 1, 1
	expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.Ints(expect)
	runSample(t, N, K, expect)
}
