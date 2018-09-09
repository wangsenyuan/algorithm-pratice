package p901

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, prices []int, expect []int) {
	spanner := Constructor()

	res := make([]int, len(prices))
	for i, price := range prices {
		res[i] = spanner.Next(price)
	}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", prices, expect, res)
	}
}

func TestSample1(t *testing.T) {
	prices := []int{100, 80, 60, 70, 60, 75, 85}
	expect := []int{1, 1, 1, 2, 1, 4, 6}
	runSample(t, prices, expect)
}
