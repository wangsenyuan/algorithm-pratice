package main

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	res := solve()

	if len(res) != N {
		t.Errorf("sample result have wrong size %d", len(res))
	}

	for i := 1; i < len(res); i++ {
		if res[i] < res[i-1] {
			t.Fatalf("Sample result fail at %d, with %d ~ %d", i, res[i], res[i-1])
		}
	}

	res = res[:4]

	expect := []int64{8, 58, 85, 88}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample result %v, not correct", res)
	}

}
