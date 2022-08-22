package p2380

import (
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := secondsToRemoveOccurrences(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "001011"
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "001001"
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "0110101"
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "11100"
	expect := 0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "01011110"
	expect := 5
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "1001111111110001011001110000000110101"
	expect := 20
	runSample(t, s, expect)
}
