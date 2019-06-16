package p736

import "testing"

func TestSample1(t *testing.T) {
	samples := []struct {
		input  string
		expect int
	}{
		{"(add 1 2)", 3},
		{"(mult 3 (add 2 3))", 15},
		{"(let x 2 (mult x 5))", 10},
		{"(let x 2 (mult x (let x 3 y 4 (add x y))))", 14},
		{"(let x 3 x 2 x)", 2},
		{"(let x 1 y 2 x (add x y) (add x y))", 5},
		{"(let x 2 (add (let x 3 (let x 4 x)) x))", 6},
		{"(let a1 3 b2 (add a1 1) b2)", 4},
		{"(let x 7 -12)", -12},
	}
	for _, sample := range samples {
		res := evaluate(sample.input)
		if res != sample.expect {
			t.Errorf("evaluate %s should give %d, but get %d", sample.input, sample.expect, res)
		}
	}
}
