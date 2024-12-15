package p3386

import (
	"math"
	"testing"
)

func runSample(t *testing.T, init string, pairs [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64, expect float64) {
	res := maxAmount(init, pairs, rates1, pairs2, rates2)

	if math.Abs(res-expect) > 1e-6 {
		t.Fatalf("Sample expect %f, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	initialCurrency := "EUR"
	pairs1 := [][]string{{"EUR", "USD"},
		{"USD", "JPY"}}
	rates1 := []float64{2.0, 3.0}
	pairs2 := [][]string{{"JPY", "USD"},
		{"USD", "CHF"},
		{"CHF", "EUR"}}
	rates2 := []float64{4.0, 5.0, 6.0}
	expect := 720.0

	runSample(t, initialCurrency, pairs1, rates1, pairs2, rates2, expect)

}
