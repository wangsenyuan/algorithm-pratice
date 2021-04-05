package b

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, logs [][]int, k int, expect []int) {
	res := findingUsersActiveMinutes(logs, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	logs := [][]int{
		{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3},
	}
	k := 5
	expect := []int{0, 2, 0, 0, 0}
	runSample(t, logs, k, expect)
}

func TestSample2(t *testing.T) {
	logs := [][]int{
		{1, 1}, {2, 2}, {2, 3},
	}
	k := 4
	expect := []int{1, 1, 0, 0}
	runSample(t, logs, k, expect)
}
