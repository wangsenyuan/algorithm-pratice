package p2750

import "testing"

func runSample(t *testing.T, num1, num2 int, expect int) {
	res := makeTheIntegerZero(num1, num2)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, -2, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 7, -1)
}

func TestSample3(t *testing.T) {
	// 34 - 18 = 16
	// num1 - x * num2 = num,
	runSample(t, 34, 9, 2)
}

func TestSample4(t *testing.T) {
	// 34 - 18 = 16
	// num1 - x * num2 = num,
	runSample(t, 85, 42, -1)
}
