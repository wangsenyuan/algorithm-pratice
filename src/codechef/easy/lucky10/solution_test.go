package main

import "testing"

func TestSample(t *testing.T) {
	samples := []struct {
		a string
		b string
		r string
	}{
		{"4", "7", "7"},
		{"7", "8", ""},
		{"435", "479", "74"},
		{"1675475", "9756417", "777744"},
	}

	for _, sample := range samples {
		res := solve([]byte(sample.a), []byte(sample.b))
		if res != sample.r {
			t.Errorf("lucky string from %s, %s should be %s, but %s", sample.a, sample.b, sample.r, res)
		}
	}
}
