package p774

import (
	"math"
	"testing"
)

func runSample(t *testing.T, stations []int, K int, expect float64) {
	res := minmaxGasDist(stations, K)
	if math.Abs(res-expect) > 0.000001 {
		t.Errorf("sample %v %d, expect %f, but got %f", stations, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	stations := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	K := 9
	expect := 0.5
	runSample(t, stations, K, expect)
}
