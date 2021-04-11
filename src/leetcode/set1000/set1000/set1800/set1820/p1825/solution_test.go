package p1825

import (
	"testing"
)

func TestSample1(t *testing.T) {
	mk := Constructor(3, 1)
	mk.AddElement(3)
	mk.AddElement(1)

	res := mk.CalculateMKAverage()

	require(t, -1, res, "res should be -1")
	mk.AddElement(10)

	res = mk.CalculateMKAverage()

	require(t, 3, res, "res should be 3")

	mk.AddElement(5)
	mk.AddElement(5)
	mk.AddElement(5)

	res = mk.CalculateMKAverage()

	require(t, 5, res, "res should be 5")
}

func require(t *testing.T, expect, res int, message string) {
	if res != expect {
		t.Fatal(message)
	}
}
