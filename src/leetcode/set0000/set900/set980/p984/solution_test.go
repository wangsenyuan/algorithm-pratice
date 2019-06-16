package p984

import "testing"

func runSample(t *testing.T, A, B int) {
	res := strWithout3a3b(A, B)

	n := len(res)

	if n != A+B {
		t.Errorf("Sample %d %d, result %s, length don't match", A, B, res)
		return
	}

	valid := true

	for i := 0; i < n; i++ {
		if i >= 2 && res[i-2] == res[i] && res[i-1] == res[i] {
			valid = false
			break
		}
		if res[i] == 'a' {
			A--
		} else {
			B--
		}
	}

	if !valid {
		t.Errorf("Sample %d %d, result %s, is not valid", A, B, res)
		return
	}

	if A != 0 || B != 0 {
		t.Errorf("Sample %d %d, result %s, is not valid", A, B, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, 2, 5)
}

func TestSample5(t *testing.T) {
	runSample(t, 2, 3)
}

func TestSample6(t *testing.T) {
	runSample(t, 27, 33)
}
