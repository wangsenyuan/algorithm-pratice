package p760

import "testing"

func runSample(t *testing.T, A, B []int) {
	res := anagramMappings(A, B)
	for i := 0; i < len(res); i++ {
		a := A[i]
		b := B[res[i]]
		if a != b {
			t.Errorf("elements at A[%d](%d) and B[%d](%d) is not same from the result %v", i, a, res[i], b, res)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{12, 28, 46, 32, 50}
	B := []int{50, 12, 32, 46, 28}
	runSample(t, A, B)
}

func TestSample2(t *testing.T) {
	A := []int{12, 28, 46, 32, 50, 46}
	B := []int{50, 12, 32, 46, 46, 28}
	runSample(t, A, B)
}
