package p374

import "testing"

func runSample(t *testing.T, n int, pick int) {
	guess = func(num int) int {
		if num > pick {
			return -1
		}
		if num < pick {
			return 1
		}
		return 0
	}
	res := guessNumber(n)

	if res != pick {
		t.Errorf("Sample expect %d, but got %d", pick, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 6)
}
